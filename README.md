# Gontainer


## How to Build

```bash
# swaggo/http-swagger 
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH
go mod tidy
swag init -g cmd/gontainer/main.go -o api/docs

# WEB
cd web/gontainer
npm run build

go build -o build/gontainer cmd/gontainer/main.go
`````