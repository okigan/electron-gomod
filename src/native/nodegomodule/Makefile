default: recompile-for-electron

recompile-for-electron: install-build-deps
	HOME=~/.electron-gyp ./node_modules/.bin/node-gyp rebuild --target=13.1.4 --arch=x64 --dist-url=https://electronjs.org/headers


.PHONY: install-build-deps
install-build-deps:
	npm install