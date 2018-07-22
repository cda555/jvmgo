package main

import "flag"
import "fmt"
import "os"

type Cmd struct { //结构体 字段的集合
	helpFlag    bool
	versionFlag bool
	cpOption    string //选项
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage //如果解析失败则调用此函数，将命令用法打印到控制台
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
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
