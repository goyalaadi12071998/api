package models

type User struct {
	ID                  int `gorm:"primaryKey"`
	Name                string
	Email               string
	Hash                string
	Salt                string
	PhoneNumber         string
	Admin               bool
	Type                string
	CountryCode         string
	EmailVerified       bool
	PhoneNumberVerified bool
	ActiveAccount       bool
	CreatedAt           int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt           int   `gorm:"autoUpdateTime:milli"`
}
