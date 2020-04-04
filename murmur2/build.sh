cd go
GOOS=js GOARCH=wasm go build -o ../main.wasm
cd ..
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .