配置go语言编译环境说明
（1）打开当前用户的环境变量配置文件  vim $HOME/.bashrc
（2）添加如下参数
export GOROOT=/home/wang/go           
export GOPATH=/home/wang/gopath      
export PATH=$PATH:$GOROOT/bin
export GOROOT_BOOTSTRAP=/home/wang/go1.4/go

交叉编译过程如下：
（1）进入golang目录，进行环境配置
cd  /usr/local/go/src/
GOOS=windows GOARCH=386 CGO_ENABLED=0 ./make.bash
（2）进入ngrok目录重新编译
cd  /usr/local/src/ngrok/
GOOS=windows GOARCH=386 make release-server release-client

（2）进入ngrok目录重新编译
cd  /usr/local/src/ngrok/
GOOS=windows GOARCH=386 make release-server release-client

//linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 ./make.bash --no-clean
CGO_ENABLED=1 GOOS=linux GOARCH=386 ./make.bash --no-clean
   
//嵌入式linux
CGO_ENABLED=1 GOOS=linux GOARCH=arm ./make.bash --no-clean
   
//嵌入式freebsd
CGO_ENABLED=1 GOOS=freebse GOARCH=amd64 ./make.bash --no-clean
CGO_ENABLED=1 GOOS=freebse GOARCH=386 ./make.bash --no-clean
//mac os
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 ./make.bash --no-clean
CGO_ENABLED=1 GOOS=darwin GOARCH=386 ./make.bash --no-clean





