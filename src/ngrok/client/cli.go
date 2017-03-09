package client

import (
	"flag"
	"fmt"
	"ngrok/version"
	"os"
)

const usage1 string = `Usage: %s [OPTIONS] <local port or address>
Options:
`

const usage2 string = `
Examples:
	ngrok 80
	ngrok -subdomain=example 8080
	ngrok -proto=tcp 22
	ngrok -hostname="example.com" -httpauth="user:password" 10.0.0.1


Advanced usage: ngrok [OPTIONS] <command> [command args] [...]
Commands:
	ngrok start [tunnel] [...]    Start tunnels by name from config file
	ngrok list                    List tunnel names from config file
	ngrok help                    Print help
	ngrok version                 Print ngrok version

Examples:
	ngrok start www api blog pubsub
	ngrok -log=stdout -config=ngrok.yml start ssh
	ngrok version

`

//选项结构体
type Options struct {
	config    string
	logto     string
	loglevel  string
	authtoken string
	httpauth  string
	hostname  string
	protocol  string
	subdomain string
	command   string
	args      []string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage1, os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage2)
	}
	//命令行 -config参数，指定客户端的配置文件，默认为$HOME/.ngrok文件
	config := flag.String(
		"config",
		"",
		"Path to ngrok configuration file. (default: $HOME/.ngrok)")
	//命令行 -log参数，指定客户端log的输出位置，默认为none（无输出）
	logto := flag.String(
		"log",
		"none",
		"Write log messages to this file. 'stdout' and 'none' have special meanings")
	//命令行 -log-level参数，指定log输出的level等级
	loglevel := flag.String(
		"log-level",
		"DEBUG",
		"The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	//命令行 -authtoken参数，指定用户的认证令牌（暂时无用后期加入用于限制用户）
	authtoken := flag.String(
		"authtoken",
		"",
		"Authentication token for identifying an ngrok.com account")
	//命令行 -httpauth参数，配置http身份认证相关参数，格式为用户名：密码，用于保护公共的隧道端点
	httpauth := flag.String(
		"httpauth",
		"",
		"username:password HTTP basic auth creds protecting the public tunnel endpoint")
	//命令行 -subdomian,配置子域名申请
	subdomain := flag.String(
		"subdomain",
		"",
		"Request a custom subdomain from the ngrok server. (HTTP only)")
	//命令行 -hostname,配置主机名申请
	hostname := flag.String(
		"hostname",
		"",
		"Request a custom hostname from the ngrok server. (HTTP only) (requires CNAME of your DNS)")
	//命令行 -proto,配置隧道的交互协议，默认http和https
	protocol := flag.String(
		"proto",
		"http+https",
		"The protocol of the traffic over the tunnel {'http', 'https', 'tcp'} (default: 'http+https')")

	flag.Parse()

	opts = &Options{
		config:    *config,
		logto:     *logto,
		loglevel:  *loglevel,
		httpauth:  *httpauth,
		subdomain: *subdomain,
		protocol:  *protocol,
		authtoken: *authtoken,
		hostname:  *hostname,
		command:   flag.Arg(0),
	}

	switch opts.command { //解析第一个参数
	case "list":
		opts.args = flag.Args()[1:] //返回非flag格式的命令行参数数组
	case "start":
		opts.args = flag.Args()[1:]
	case "version":
		fmt.Println(version.MajorMinor())
		os.Exit(0)
	case "help":
		flag.Usage()
		os.Exit(0)
	case "": //第一个参数为空输出错误
		err = fmt.Errorf("Error: Specify a local port to tunnel to, or " +
			"an ngrok command.\n\nExample: To expose port 80, run " +
			"'ngrok 80'")
		return

	default:
		if len(flag.Args()) > 1 { //如果非flag格式的参数个数超过1个，输出错误
			err = fmt.Errorf("You may only specify one port to tunnel to on the command line, got %d: %v",
				len(flag.Args()),
				flag.Args())
			return
		}

		opts.command = "default"
		opts.args = flag.Args()
	}

	return
}
