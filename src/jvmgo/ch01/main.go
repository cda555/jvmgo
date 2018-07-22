package main

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag { //如果 输入 -version
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" { //如果用户输入 -help 或者解析出错则
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Println(cmd.cpOption)
	fmt.Printf("classpath: %s class: %s args: %v \n",
		cmd.cpOption, cmd.class, cmd.args)
}
