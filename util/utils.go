package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// GenerateFlag To generate keys for tests
func GenerateFlag(name string) string {
	// Generate use md5 with plain text in format "TEST_NAME + TIMESTAMP + COSTUMER_KEY"
	byteMd5 := md5.Sum([]byte(fmt.Sprintf(name, time.Now().String(), "hE.eA.aD.dE.eR.rF.fA.aK.kE.eR.rH")))
	return fmt.Sprint("flag{", strings.ToLower(hex.EncodeToString(byteMd5[:])), "}")
}
