#!/bin/sh
rm -rfv pkg/linux_amd64/ngrok/ 
rm -fv bin/ngrok bin/ngrokd src/ngrok/client/assets/*.go src/ngrok/server/assets/*.go
rm -rfv bin/linux_amd64/
GOOS=linux GOARCH=amd64 make release-server release-client
