package constants

import "time"

const (
	JWT_ACCESS_TOKEN_TIME_DURATION     = time.Duration(time.Minute * 60)
	JWT_REFRESH_TOKEN_TIME_DURATION    = time.Duration(time.Hour * 24)
	TTL_COOKIE_EXPIRY_SESSION_DURATION = time.Duration(time.Minute * 60)
	TTL_FOR_RATE_LIMIT                 = time.Second * 60
	MAX_REQUEST_ALLOWED_PER_MIN        = 10
)
