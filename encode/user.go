package encode

// User struct
type User struct {
	UID    uint32  `json:"uid"`
	Nick   string  `json:"nick"`
	Score  float32 `json:"score"`
	Gender uint8   `json:"gender"`
	Age    uint8   `json:"age"`
}
