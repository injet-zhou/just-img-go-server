package config

const (
	DEV    = "dev"
	PROD   = "prod"
	ENVkEY = "JUST_IMG_GO_ENV"
)

const (
	PORT = "7780"
)

type PlatformType int

const (
	OSS PlatformType = iota + 1
	COS
	QINIU
	UPYUN
	Local
)

const (
	MAX_LOGIN_FAIL_COUNT = 5
	USER_SESSION_KEY     = "USK_"
	TOKEN_EXPIRE_TIME    = 3 * 60 * 60
)
