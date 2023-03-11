package model

type Book struct {
	ID           int    `gorm:"primaryKey" json:"id"`
	Stage        int    `json:"stage"`
	Class        string `json:"class"`
	Descriptions string `json:"descriptions"`
	Users        string `json:"users"`
}
