test:
	go test . -v
coverage:
	go test -coverprofile=coverage.out
coverage-html:
	make coverage && go tool cover -html=coverage.out
test-suite:
	go test -v -run $(suite)