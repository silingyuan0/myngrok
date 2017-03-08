#!/bin/sh
rm -rfv pkg/windows_amd64/ngrok/ 
rm -fv  src/ngrok/client/assets/*.go src/ngrok/server/assets/*.go
rm -rfv bin/windows_amd64/
GOOS=windows GOARCH=amd64 make release-server release-client
