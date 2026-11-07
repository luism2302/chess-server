run:
	go run .

test/run_tests:
	go test ./... -v
test/coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

