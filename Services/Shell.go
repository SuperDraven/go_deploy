package Services

import (
	"fmt"
	"os/exec"
)

func PullShellGo(directory string, c string)  {
	c = "cd " +directory+" && "+c
	fmt.Println(c)
	cmd := exec.Command("/bin/bash", "-c", c)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
	//return
}