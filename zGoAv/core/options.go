package core

import (
	"flag"
	"github.com/zGoAv/gologger"
	"os"
)

type Options struct {
	/*
		FileName shellcode文件名
		Manual shellocde生成方法
	*/
	FileName string
	Means    bool
	Stdin    bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.StringVar(&options.FileName, "f", "", "通过Shellcode生成免杀马")
	flag.BoolVar(&options.Means, "means", false, "查看Shellcode生成方法")
	flag.Parse()
	options.Stdin = hasStdin()
	//ShowBanner()
	/*
			&& 逻辑运算符AND，如果两边操作都是True，则条件为True，否则为False
			! 逻辑运算符NOT，如果条件为True，则逻辑NOT条件为False，否则为True
		代码实例：
	*/
	//package main
	//
	//import "fmt"
	//
	//func main() {
	//	var a bool = true
	//	var b bool = false
	//	if ( a && b ) {
	//		fmt.Printf("第一行 - 条件为 true\n" )
	//	}
	//	if ( a || b ) {
	//		fmt.Printf("第二行 - 条件为 true\n" )
	//	}
	//	/* 修改 a 和 b 的值 */
	//	a = false
	//	b = true
	//	if ( a && b ) {
	//		fmt.Printf("第三行 - 条件为 true\n" )
	//	} else {
	//		fmt.Printf("第三行 - 条件为 false\n" )
	//	}
	//	if ( !(a && b) ) {
	//		fmt.Printf("第四行 - 条件为 true\n" )
	//	}
	//}
	if len(options.FileName) == 0 && !options.Means {
		flag.Usage()
		os.Exit(0)
	}
	if options.FileName != "" && !FileExists(options.FileName) {
		gologger.Fatalf("文件%s 不存在！\n", options.FileName)
		os.Exit(0)
	}
	return options
}

func hasStdin() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false
	}
	return true
}
