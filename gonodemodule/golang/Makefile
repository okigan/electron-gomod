default: compile-go

compile-go:
	go build -buildmode=c-archive -o markdown.so golang.go

recompile-for-electron:
	HOME=~/.electron-gyp ./node_modules/.bin/node-gyp rebuild --target=13.1.4 --arch=x64 --dist-url=https://electronjs.org/headers

build: compile-go
	npm run build