package Services

import (
	"fmt"
	"os/exec"
)

func ExecuteShell(c string)  {
	cmd := exec.Command("/bin/bash", "-c", c)
	fmt.Println(c)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
func  PullShellTest(shell string)  {
	command := `./Shell/`+shell+`.sh`
	cmd := exec.Command("/bin/bash", "-C", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}
