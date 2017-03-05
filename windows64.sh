#!/bin/sh
rm -rfv pkg/windows_amd64/ngrok/ src/ngrok/client/assets/ src/ngrok/server/assets/ bin/windows_386/
GOOS=windows GOARCH=amd64 make release-server release-client
