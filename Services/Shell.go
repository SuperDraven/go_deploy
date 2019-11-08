package Services

import (
	"fmt"
	"os/exec"
)

func ShellGo(c string)  {
	cmd := exec.Command("/bin/bash", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}