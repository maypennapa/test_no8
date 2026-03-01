package models

type Question struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Number   int    `json:"number"`
	Question string `json:"question"`
	Choice1  string `json:"choice1"`
	Choice2  string `json:"choice2"`
	Choice3  string `json:"choice3"`
	Choice4  string `json:"choice4"`
}
