include .env
export

.PHONY: test format test32

.ONESHELL:

test32:
	@go test \
		-count=1 \
		-v ./tests/medium/reverse_integer... | sed -e "/PASS/s//$$(printf "\033[32mPASS\033[0m")/" \
			-e "/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/" \
			-e "/SKIP/s//$$(printf "\033[33mSKIP\033[0m")/" \
			-e "/^=== RUN/d" \
			-e "/^\?/d"

test:
	@go test \
		-count=1 \
		-v ./tests/medium/sort_colors/... | sed -e "/PASS/s//$$(printf "\033[32mPASS\033[0m")/" \
			-e "/FAIL/s//$$(printf "\033[31mFAIL\033[0m")/" \
			-e "/SKIP/s//$$(printf "\033[33mSKIP\033[0m")/" \
			-e "/^=== RUN/d" \
			-e "/^\?/d"
format:
	@go fmt