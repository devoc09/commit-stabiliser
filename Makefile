# build scripts
# build binary for Mac
darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/cst ./main.go

# build binary for Linux AMD64 architecture
linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux64/cst ./main.go

# build binary for Linux ARM architecture
linuxArm:
	GOOS=linux GOARCH=arm go build -o bin/linuxArm/cst ./main.go
