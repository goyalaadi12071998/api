package constants

import "time"

const (
	JWT_ACCESS_TOKEN_TIME_DURATION  = time.Duration(time.Minute * 60)
	JWT_REFRESH_TOKEN_TIME_DURATION = time.Duration(time.Hour * 24)
)
