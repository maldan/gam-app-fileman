GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o app .
cp app /root/.gam-app/maldan-gam-app-fileman-v0.0.15