#!/bin/sh
rm -rfv pkg/windows_386/ngrok/ 
rm -fv  src/ngrok/client/assets/*.go src/ngrok/server/assets/*.go
rm -rfv bin/windows_386/

GOOS=windows GOARCH=386 make release-server release-client
#GOOS=windows GOARCH=386 make
