package utils

import "time"

// GetCurrTS return current timestamps
func GetCurrTS() int64 {
	return time.Now().Unix()
}
