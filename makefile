build:
	@go run dist.go

test:
	@go run devserver.go

clean:
	@-rm main_wasm.go
	@-rm *_vgen.go
	@-rm -r dist
