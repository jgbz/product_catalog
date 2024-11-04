.PHONY: up down run_integration teardown_integration run_tests

up:
	docker compose -f ./build/docker-compose.yaml up --build

down:
	docker compose -f ./build/docker-compose.yaml down --volumes

teardown_integration:
	docker compose -f ./build/docker-compose-test.yaml down --volumes

run_integration: teardown_integration
	docker compose -f ./build/docker-compose-test.yaml up --build

run_tests:
	go test ./... -short -count=1