build:
	@go run setup.go
	@vugugen -s -r src
	cd src && GOOS=js GOARCH=wasm go build -o main.wasm
	@mv src/main.wasm dist/main.wasm
	@make cleanbuild

cleanbuild:
	@-rm src/main_wasm.go
	@-rm src/*_vgen.go

clean:
	@make cleanbuild
	@-rm -r dist

test:
	@go run devserver.go

