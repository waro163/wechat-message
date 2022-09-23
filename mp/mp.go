package mp

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func Validate(req *http.Request, token string, skip bool) bool {
	if skip {
		return true
	}
	query := req.URL.Query()
	timestamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	signature := query.Get("signature")
	return signature == Signature(token, timestamp, nonce)
}

func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
