build:
	mkdir -p build/web/resources && cp -r resources build/web
	GOARCH=wasm GOOS=js go build -o ./build/web/app.wasm main.go
	go build -o ./build/app main.go

run:
	@if lsof -t -i :8080; \
    then \
        kill -9 $$(lsof -t -i:8080); \
    fi
	
	cd build && ./app


install-tailwind:
	wget https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 $(GOPATH)/bin/tailwindcss

build-tailwind:
	cd tailwind && tailwindcss -i ./tailwind-min.css -o ../resources/styles/tailwind/tailwind-min.css --minify

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