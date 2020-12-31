package udplogs

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	// "sync"
	"time"
)

const (
	ListenPort           = ":13333"
	LogFileTimeOut int64 = 50
	CheckInterval        = 30
)

type udpParam struct {
	conn *net.UDPConn
	f    *os.File
}

type LogFile struct {
	f        *os.File
	LastTime int64
	// mu *sync.Mutex
}

// type LogFileMap map[string]*LogFile

var AllLogFile = make(map[string]*LogFile)

type LogData struct {
	FileName string `json:"filename"`
	Content  string `json:"content"`
	Time     int    `json:"time"`
}

func Listen() {
	f, _ := os.OpenFile("udplogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()

	fmt.Println("udp logs server start...")
	udpAddr, err := net.ResolveUDPAddr("udp4", ListenPort)
	if err != nil {
		panic(err)
	}

	listen, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	up := &udpParam{listen, f}

	go up.CheckLogFile()

	for {
		up.handleUDPConection()
	}
}

func (up *udpParam) CheckLogFile() {
	i := 0
	for range time.Tick(time.Second * CheckInterval) {
		now := time.Now()
		timeStr := getTimeStr(now)
		fmt.Println(timeStr, "check file...", i)
		i++
		for k, v := range AllLogFile {
			fmt.Println(fmt.Sprintf("%s time: %d", k, v.LastTime))
			if now.Unix() > v.LastTime+LogFileTimeOut {
				v.f.Close()
				up.f.WriteString(fmt.Sprintf("%s|logfile %s close!\n", timeStr, k))
				delete(AllLogFile, k)
			}
		}
	}
}

func (up *udpParam) handleUDPConection() {
	buf := make([]byte, 4096)
	n, addr, err := up.conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("UDP client: ", addr)
	fmt.Println("Get messages: ", string(buf[:n]))

	up.conn.WriteToUDP([]byte(time.Now().String()), addr)
	go up.WriteLog(addr.String(), buf[:n])
}

func (up *udpParam) WriteLog(ip string, logBytes []byte) bool {
	if len(logBytes) == 0 {
		return false
	}
	logdata := &LogData{}
	err := json.Unmarshal(logBytes, logdata)
	if err != nil {
		return false
	}

	filename := logdata.FileName
	now := time.Now()
	if _, ok := AllLogFile[filename]; !ok {
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return false
		}
		AllLogFile[filename] = &LogFile{f, now.Unix()}
		up.f.WriteString(fmt.Sprintf("%s|create file object: %s\n", getTimeStr(now), filename))
	}

	AllLogFile[filename].LastTime = now.Unix()
	logStr := fmt.Sprintf("%s|%s|%s", getTimeStr(now), ip, logdata.Content)
	fmt.Println(logStr)
	AllLogFile[filename].f.WriteString(logStr + "\n")
	return true
}

func getTimeStr(time time.Time) string {
	return time.Format("2006-01-02 15:04:05")
}
