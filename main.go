package main

import (
	"fmt"
	"os/exec"
)


func main()  {
	c := "git pull"
	cmd := exec.Command("/bin/bash", "-c", c)
	out, err := cmd.Output()
	fmt.Println(err)
	fmt.Println(string(out))
}
