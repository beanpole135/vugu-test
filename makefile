build:
	@go run setup.go
	@vugugen -s -r src
	cd src && GOOS=js GOARCH=wasm go build -o main.wasm
	@mv src/main.wasm dist/main.wasm

test:
	@go run devserver.go

clean:
	@-rm main_wasm.go
	@-rm *_vgen.go
	@-rm -r dist
