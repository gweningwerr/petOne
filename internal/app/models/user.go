package models

//type User struct {
//	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
//	Login    string `gorm:"size:255;not null" json:"login_name"`
//	Password string `gorm:"size:255;not null" json:"password"`
//}

type User struct {
	ID          uint64  `json:"id"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	PhoneNumber *string `json:"phone_number"`
}
