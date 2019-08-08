package utils

import (
	"fmt"
	"time"
)

//11为时间戳
func DtToString(dt int64) string {
	t := time.Unix(dt, 0)
	return fmt.Sprintf("%d/%d/%d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
