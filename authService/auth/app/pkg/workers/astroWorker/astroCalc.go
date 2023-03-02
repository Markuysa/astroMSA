package astroWorker

import (
	"fmt"
	"time"
)

func CalculateSign(birthTime time.Time) string {

	birthDay := birthTime.Day()
	birthMonth := birthTime.Month()
	fmt.Println(birthDay, birthMonth)
	return ""
}
