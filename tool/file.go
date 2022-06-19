package tool

import (
	"fmt"
	"time"
)

func DefaultPath() string {
	now := time.Now()
	return fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())
}
