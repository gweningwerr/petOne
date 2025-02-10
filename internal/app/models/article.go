package models

//type Article struct {
//	ID      uint64 `gorm:"primary_key"`
//	Title   string `gorm:"type:varchar(255);not null"`
//	Author  string `gorm:"type:varchar(255);not null"`
//	Content string `gorm:"type:text;not null"`
//}

type Article struct {
	ID      uint64 `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
