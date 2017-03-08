#!/bin/sh
rm -rfv pkg/linux_386/ngrok/ 
rm -fv bin/ngrok bin/ngrokd src/ngrok/client/assets/*.go src/ngrok/server/assets/*.go
rm -rfv bin/linux_386/
#GOOS=linux GOARCH=386 make release-server release-client
GOOS=linux GOARCH=386 make

