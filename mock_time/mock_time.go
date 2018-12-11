package mock_time

import (
	"fmt"
	"os/exec"
)

func SetTime(year int, month int, day int) {
	str := getMacDateTimeString(year, month, day)
	exec.Command("sudo", "date", str).Run()
}

func getMacDateTimeString(year int, month int, day int) string {
	return fmt.Sprint("%d%d0000%d", month, day, year)
}