package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// GetMd5String Generate a plain string's md5
func GetMd5String(plain string) string {
	byteMd5 := md5.Sum([]byte(plain))
	return hex.EncodeToString(byteMd5[:])
}

// GenerateFlag To generate keys for tests
func GenerateFlag(name string) string {
	// Generate use md5 with plain text in format "TEST_NAME + TIMESTAMP + COSTUMER_KEY"
	md5Flag := GetMd5String(fmt.Sprintf(name, time.Now().String(), "hE.eA.aD.dE.eR.rF.fA.aK.kE.eR.rH"))
	return fmt.Sprint("flag{", strings.ToLower(md5Flag), "}")
}
