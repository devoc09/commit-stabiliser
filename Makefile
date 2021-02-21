# Setup scripts
# build binary for Mac and set command into your GOPATH/bin
darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/git-cst ./main.go
	mv bin/mac/git-cst ${GOPATH}/bin/

# build binary for Linux AMD64 architecture and set command into your GOPATH/bin
linux64:
	GOOS=linux GOARCH=amd64 go build -o bin/linux64/git-cst ./main.go
	mv bin/linux64/git-cst ${GOPATH}/bin/

# build binary for Linux ARM architecture and set command into your GOPATH/bin
linuxArm:
	GOOS=linux GOARCH=arm go build -o bin/linuxArm/git-cst ./main.go
	mv bin/linuxArm/git-cst ${GOPATH}/bin/

# Only build binary for darwin, linux64, and linuxArm
build:
	GOOS=darwin GOARCH=amd64 go build -o bin/mac/git-cst ./main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux64/git-cst ./main.go
	GOOS=linux GOARCH=arm go build -o bin/linuxArm/git-cst ./main.go

