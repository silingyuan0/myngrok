#!/bin/sh
rm -rfv pkg/linux_amd64/ngrok/ src/ngrok/client/assets/ src/ngrok/server/assets/ 
rm -fv bin/ngrok bin/ngrokd
GOOS=linux GOARCH=amd64 make release-server release-client
