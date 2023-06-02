test:
	go test . -v
cov:
	go test -coverprofile=coverage.out
cov-html:
	make cov && go tool cover -html=coverage.out
test-suite:
	go test -v -run $(suite)