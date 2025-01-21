include .env
export

.PHONY: test format

.ONESHELL:

test:
	@go test \
		-count=1 \
		-v ./... | sed -e "/PASS/s//$$(printf "\033[32mPASS\033[0m")/" \
			-e "/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/" \
			-e "/SKIP/s//$$(printf "\033[33mSKIP\033[0m")/" \
			-e "/^=== RUN/d" \
			-e "/^\?/d"
format:
	@go fmt