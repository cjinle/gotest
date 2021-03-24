package cjson

import (
	"encoding/json"
	"log"
)

// Message 推送内容
type Message struct {
	API       int      `json:"api"`
	Tokens    []string `json:"tokens"`
	Content   string   `json:"content"`
	StatKey   string   `json:"statkey,omitempty"`
	Sandbox   int      `json:"sandbox,omitempty"`
	ID        string   `json:"id,omitempty"`
	LoginType string   `json:"loginType,omitempty"`
}

// Data struct
type Data struct {
	Content   string `json:"content"`
	ID        string `json:"id,omitempty"`
	LoginType string `json:"loginType,omitempty"`
}

// Payload struct
type Payload struct {
	RegistrationIDs []string `json:"registration_ids"`
	Data            Data     `json:"data"`
}

// Encode out put string
func Encode() string {

	jsonStr := `{"api":123123,"tokens":["123","456","789"],"content":"hello world","ID":"xxx"}`
	var msg Message
	json.Unmarshal([]byte(jsonStr), &msg)
	log.Println(msg)
	payload := Payload{
		RegistrationIDs: msg.Tokens,
		Data:            Data{Content: msg.Content, ID: msg.ID, LoginType: msg.LoginType},
	}
	log.Println(payload)
	s, _ := json.Marshal(payload)
	return string(s)
}
