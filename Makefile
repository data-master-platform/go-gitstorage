
.PHONY: test-in-docker
test-in-docker:
	docker-compose up --build --abort-on-container-exit && docker-compose rm -fsv

.PHONY: test-all
test-all:
	go test ./... -v -coverprofile=coverage/coverage.out -covermode=atomic

.PHONY: test
test:
	go test -short ./... -v

.PHONY: test-integration
test-integration:
	go test -run  TestIntegration* ./... -v