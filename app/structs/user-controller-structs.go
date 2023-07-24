package structs

type UserSingupRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CountryCode string `json:"country_code"`
	Type        string `json:"type"`
}

type UserSingupRequestResponse struct {
	Id int `json:"id"`
}
