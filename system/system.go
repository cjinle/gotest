package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Exec() {
	log.Println("exec")
	cmd := exec.Command("ls", "-lh")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(stdout))
}

func ExecAsync() {
	picUrl := "https://avatar.csdnimg.cn/7/8/E/3_butterfly5211314.jpg"
	cmd := exec.Command("wget", picUrl)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func ExecAsync2() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	videoUrl := "http://pgc.qcdn.xiaodutv.com/2112464083_732621686_2020120716041420201207164242.mp4?Cache-Control=max-age%3D8640000&responseExpires=Wed%2C+17+Mar+2021+16%3A45%3A12+GMT&xcode=aea3dc193494fe118a57cf801729de96af6c7348b9427965&time=1607514739"
	videoTitle := "孟鹤堂周九良人气太高，还没出场就被全场观众高呼，盘他！.mp4"
	videoTitle = "xxx.mp4"
	cmd := exec.Command("ffmpeg", "-i", `"`+videoUrl+`"`, "-c", "copy", `"`+videoTitle+`"`)
	log.Println(cmd.Path, cmd.Args)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err, stderr.String())
	}
	log.Println(out.String())
}

func MacAddr() {
	fmt.Println(net.InterfaceAddrs())
	inters, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range inters {
		log.Println(v.Name, v.HardwareAddr)
	}
}

func main() {
	MacAddr()
}
