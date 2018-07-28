help: ##@other Show this help
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

install: ##@dev Installs current package
	go install ./...

serve: ##@dev Starts the server to process api requests
	go run main.go assign_handler.go section_handler.go get_seat_handler.go

test: ##@test Runs associated tests
	go get -u github.com/gorilla/mux
	go get -u github.com/stretchr/testify/assert
	go test ./...

add-section: ##@dev Calls endpoint to add a new section using curl
	curl -X POST --data '{"id":"mysection","rows":1,"blocks":[3,4,3]}' localhost:9000/sections

assign-seat: ##@dev Calls endpoint to assign a seat to default section
	curl -X POST --data '{"id":"mysection"}' localhost:9000/sections/assign

get-seat-number: ##@dev Get the assigned seat number (e.g., 32C, 44F, etc) given the index of the seat
	curl -X GET localhost:9000/sections/seat?number=$(number)\&section=mysection

lint-install: ##@lint installs necessary packages to run linting tools
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b $(GOPATH)/bin v1.9.1

lint: ##@lint runs the linter
	@golangci-lint run ./...


# This is a code for automatic help generator.
# It supports ANSI colors and categories.
# To add new item into help output, simply add comments
# starting with '##'. To add category, use @category.
GREEN  := $(shell echo "\e[32m")
WHITE  := $(shell echo "\e[37m")
YELLOW := $(shell echo "\e[33m")
RESET  := $(shell echo "\e[0m")

HELP_FUN = \
		   %help; \
		   while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z0-9\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
		   print "Usage: make [target]\n\n"; \
		   for (sort keys %help) { \
			   print "${WHITE}$$_:${RESET}\n"; \
			   for (@{$$help{$$_}}) { \
				   $$sep = " " x (32 - length $$_->[0]); \
				   print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
			   }; \
			   print "\n"; \
		   }
