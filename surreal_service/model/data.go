package model

import (
	"time"
)

type Message struct {
	UserID    string `json:"id,omitempty"`
	SessionID string `json:"session_id"`
	Content   string `json:"content"`
	MsgNum    int    `json:"msg_num"`
}

type PluginNode struct {
	ID      string `json:"id,omitempty"`
	Tipo    int    `json:"tipo"`
	Content string `json:"content"`
}

type Relation struct {
	From string `json:"from"`
	To   string `json:"to"`
	Type string `json:"type"` // "temporale" o "dipendenza o aggiunge "
}

type LastMsg struct {
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

type LastMsgContainer struct {
	ID       string    `json:"id"`
	Messages []LastMsg `json:"messages"`
}

type AddMessageRequest struct {
	ContainerID string `json:"containerID"`
	Text        string `json:"text"`
}
