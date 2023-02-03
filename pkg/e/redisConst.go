package e

import "time"

const (
	VerificationCodeKey    = "lovenature:user:code:"
	VerificationCodeKeyTTL = time.Second * 60 * 3

	UserLoginInfo    = "lovenature:user:token:"
	UserLoginInfoTTL = time.Hour * 24 * 30

	PetHotData    = "lovenature:hotData:pet:"
	PetHotDataDDL = time.Second * 60
)
