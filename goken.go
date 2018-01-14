package goken

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func GenerateToken(key, userID, actionID string) string {
	if len(key) == 0 {
		panic("zero length xsrf secret key")
	}
	// Round time up and convert to milliseconds.
	milliTime := strconv.FormatInt((time.Now().UnixNano()+1e6-1)/1e6, 10)

	h := hmac.New(sha1.New, *(*[]byte)(unsafe.Pointer(&key)))
	m := strings.Replace(userID, ":", "_", -1) + ":" + strings.Replace(actionID, ":", "_", -1) + ":" + milliTime

	h.Write(*(*[]byte)(unsafe.Pointer(&m)))

	src := h.Sum(nil)
	enc := base64.URLEncoding
	dest := make([]byte, enc.EncodedLen(len(src)))
	enc.Encode(dest, src)

	return strings.TrimRight(*(*string)(unsafe.Pointer(&dest)), "=") + ":" + milliTime
}

// generateTokenAtTime is like Generate, but returns a token that expires 24 hours from now.
func GenerateTokenAtTime(key, userID, actionID string) string {
	if len(key) == 0 {
		panic("zero length xsrf secret key")
	}
	// Round time up and convert to milliseconds.
	milliTime := (time.Now().UnixNano() + 1e6 - 1) / 1e6

	h := hmac.New(sha1.New, []byte(key))
	fmt.Fprintf(h, "%s:%s:%d", strings.Replace(userID, ":", "_", -1), strings.Replace(actionID, ":", "_", -1), milliTime)

	// Get the padded base64 string then removing the padding.
	tok := string(h.Sum(nil))
	tok = base64.URLEncoding.EncodeToString([]byte(tok))
	tok = strings.TrimRight(tok, "=")

	return fmt.Sprintf("%s:%d", tok, milliTime)
}
