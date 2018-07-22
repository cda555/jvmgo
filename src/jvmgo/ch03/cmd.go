package main

import "flag"
import "fmt"
import "os"

//Cmd  参数
type Cmd struct { //结构体 字段的集合
	helpFlag    bool
	versionFlag bool
	cpOption    string //类路径参数
	XjreOption  string //jre 路径
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage //如果解析失败则调用此函数，将命令用法打印到控制台
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath") //类路径
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")

	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre") //jre路径

	flag.Parse() //解析
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0] // 第一个参数 主类名
		cmd.args = args[1:] // 传给主类的参数
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
