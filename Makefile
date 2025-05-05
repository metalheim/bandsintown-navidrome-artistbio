all: bandsintown-navidrome-artistbio

bandsintown-navidrome-artistbio: plugin.wasm

plugin.wasm: plugin.go
	GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o $@ ./

clean:
	rm -f plugin.wasm