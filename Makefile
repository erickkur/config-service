dep: ## Get the dependencies
	@go get -v ./...

run:
	source local-env.sh && ./tools/run.sh

test:
	go test -cover ./...