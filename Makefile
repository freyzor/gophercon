PROJECT?=github.com/freyzor/gophercon
APP?=gophercon
PORT?=8000

clean:
	rm -f ./bin/${APP}

build: clean
	#CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH}
	go build \
 		-o /bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	PORT=${PORT} ./bin/${APP}

test:
	go test -race ./..
