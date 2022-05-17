package pkg

import (
	"fmt"
	"time"
)

func GetDatetime() string {
	fmt.Print(time.Now())
	timenow := time.Now().Format(time.RFC3339)
	date := fmt.Sprintf("%v\n", timenow)
	return date
}
