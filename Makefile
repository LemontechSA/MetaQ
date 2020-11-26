# LOCAL

clean:
	- rm -rf coverage
	- rm -rf tmp
	- rm -rf vendor
	- rm main
	- rm metaq

install:
	go mod download

run:
	go run ./cmd/metaq/main.go

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o metaq ./cmd/metaq/main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o metaq ./cmd/metaq/main.go

####################################

# SUPPORT

mysql:
	docker run --rm -it \
	-p 3306:3306 \
	--name mysql \
	-e MYSQL_ROOT_PASSWORD=password \
	mysql:5.6
