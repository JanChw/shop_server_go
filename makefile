dev:
	ENV=development gin -i -b server run server.go
prod:
	GIN_MODE=release ENV=production go build server.go