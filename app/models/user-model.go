package models

type User struct {
	Id                  int `gorm:"primary_key;auto_increment"`
	Name                string
	Email               string
	PhoneNumber         string
	Hash                string
	Salt                string
	CountryCode         string
	EmailVerified       bool
	PhoneNumberVerified bool
	ActiveAccount       bool
	CreatedAt           string `gorm:"autoCreateTime:milli"`
	UpdatedAt           string `gorm:"autoCreateTime:milli"`
}
