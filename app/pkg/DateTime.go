package pkg

import (
	"fmt"
	"time"
)

func GetDatetime() string {
	timenow := time.Now().UTC().Format("2006-01-02 15:04:05")
	date := fmt.Sprintf("%v\n", timenow)
	return date
}
