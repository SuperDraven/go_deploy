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
func  PullShellTest(shell string)  {
	command := `./Shell/`+shell+`.sh`
	cmd := exec.Command("/bin/bash", "-C", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
