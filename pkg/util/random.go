package util

import (
	"github.com/google/uuid"
	"lovenature/log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func RandomCode(len int) string {
	rand.Seed(time.Now().UnixNano())
	code := strings.Builder{}
	for i := 0; i < len; i++ {
		code.WriteString(strconv.Itoa(rand.Intn(10)))
	}
	return code.String()
}

func NextToken() string {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Error(err)
		return ""
	}
	return u.String()
}
