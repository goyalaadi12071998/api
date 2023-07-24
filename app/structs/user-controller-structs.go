package structs

type UserSingupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	Type        string `json:"type"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSingupRequestResponse struct {
	Id int `json:"id"`
}

type UserLoginRequestResponse struct {
	Id                  int    `json:"id"`
	Name                string `json:"name"`
	Email               string `json:"email"`
	CountryCode         string `json:"country_code"`
	Type                string `json:"type"`
	PhoneNumber         string `json:"phone_number"`
	Admin               bool   `json:"admin"`
	EmailVerified       bool   `json:"email_verified"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	ActiveAccount       bool   `json:"active_account"`
	CreatedAt           int64  `json:"created_at"`
	UpdatedAt           int    `json:"updated_at"`
}
