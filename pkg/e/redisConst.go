package e

import "time"

const (
	VerificationCodeKey    = "lovenature:user:code:"
	VerificationCodeKeyTTL = time.Second * 60

	UserLoginToken    = "lovenature:user:token:"
	USerLoginTokenTTL = time.Hour * 24 * 30
)
