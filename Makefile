AOC_DAYS := $(shell find . -maxdepth 1 -type d -name 'day*')
LATEST_AOC_DAY := $(lastword $(AOC_DAYS))

.PHONY: all
all: $(LATEST_AOC_DAY)

.PHONY: $(AOC_DAYS)
$(AOC_DAYS):
	go run ./$@

.PHONY: test
test:
	go test ./day*
