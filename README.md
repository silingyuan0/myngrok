运行pre.sh可以生成ngrok的数字签名
注释：在实际部署过程中在虚拟机linux下编译的ngrok程序再服务器中不能运行
如果在服务器中运行ngrokd出现错误时，可在服务器中搭建go编译环境重新编译ngrok源文件即可
切记：服务器中的key和客户端内的key要保持一致否则会出现key验证失败的情况，即将服务器中的签名文件复制到客户端机器中再编译客户端程序。
运行windows.sh可生成windows 32位的应用程序
运行linux.sh可生成linux 32位的应用程序
