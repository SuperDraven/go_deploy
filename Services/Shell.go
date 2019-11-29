package Services

import (
	"bufio"
	"fmt"
	"io"
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
	commandName := "/bin/bash"

	params :=[]string{"-c",  shell}
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}

	_ = cmd.Start()
	//fmt.Println(errrs)
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		//fmt.Println(err2)
		fmt.Println(line)
	}

	_ = cmd.Wait()
}
