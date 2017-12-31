package mock_time

import (
	"os/exec"
)

func SetTime() {
	exec.Command("sudo", "date", "0101000018").Run()
}