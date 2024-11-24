# Gontainer


## How to Build

```bash
# swaggo/http-swagger 
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/gontainer/main.go -o api/docs

# WEB
cd web/gontainer
npm run build

go build -o build/gontainer cmd/gontainer/main.go
```