package client

import (
	"fmt"
	"math/rand"
	"ngrok/log"
	"ngrok/util"
	"os"
	"runtime"
	"time"

	"github.com/inconshreveable/mousetrap"
)

func init() {
	if runtime.GOOS == "windows" { //如果在windows平台下使用双击方式打开则退出
		if mousetrap.StartedByExplorer() {
			fmt.Println("Don't double-click ngrok!")
			fmt.Println("You need to open cmd.exe and run it from the command line!")
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}
	}
}

func Main() {
	// 解析输入参数
	opts, err := ParseArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 设置调试输出
	log.LogTo(opts.logto, opts.loglevel)

	// 读取配置文件
	config, err := LoadConfiguration(opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 产生随机种子数
	seed, err := util.RandomSeed()
	if err != nil {
		fmt.Printf("Couldn't securely seed the random number generator!")
		os.Exit(1)
	}
	rand.Seed(seed)

	NewController().Run(config)
}
