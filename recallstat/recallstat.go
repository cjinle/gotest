package recallstat

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"encoding/json"
)

type Msg struct {
	MulticastId  int64    `json:"multicast_id"`
	Success      int      `json:"success"`
	Failure      int      `json:"failure"`
	CanonicalIds int      `json:"canonical_ids"`
	Results      []Result `json:"results"`
}

type Result struct {
	MessageId string `json:"message_id"`
	Error     string `json:"error"`
}

func Decode(jsonStr string) (*Msg, error) {
	msg := &Msg{}
	if len(jsonStr) == 0 {
		return msg, fmt.Errorf("json data empty")
	}
	err := json.Unmarshal([]byte(jsonStr), msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

type StatData map[string]int

func (sd *StatData) Stat(str string) error {
	msg, err := Decode(str)
	if err != nil {
		return err
	}
	for _, v := range msg.Results {
		if _, ok := (*sd)[v.Error]; !ok {
			(*sd)[v.Error] = 1
		} else {
			(*sd)[v.Error] += 1
		}
	}
	return nil
}

func ReadFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	defer file.Close()
	sd := make(StatData)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "|")
		err = sd.Stat(arr[1])
		if err != nil {
			fmt.Println(err)
		}
	}
	return fmt.Sprintf("%v", sd)
}
