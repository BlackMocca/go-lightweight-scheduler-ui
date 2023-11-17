build:
	mkdir -p build/web/styles && cp -R styles build/web
	GOARCH=wasm GOOS=js go build -o ./build/web/app.wasm main.go
	go build -o ./build/app main.go

run:
	@if lsof -t -i :8080; \
    then \
        kill -9 $$(lsof -t -i:8080); \
    fi
	
	cd build && ./app

clean:
	go clean ./...

mockery:
	mockery --all --dir service/$(service) --output service/$(service)/mocks

test:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o coverage.html
	export unit_total=$$(go test ./... -v  | grep -c RUN) && echo "Unit Test Total: $$unit_total" && export coverage_total=$$(go tool cover -func cover.out | grep total | awk '{print $$3}') && echo "Coverage Total: $$coverage_total"

genproto:
	protoc -I . --go_out=plugins=grpc:proto/ proto/*.proto

docker-build-prod:
	docker build -f Dockerfile-production -t godflow-ui . 