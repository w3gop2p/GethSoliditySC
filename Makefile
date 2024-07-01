build:
	@go build  -o bin/GethSoliditySC
run: build
	@./bin/GethSoliditySC
test :
	@go test -v ./...

