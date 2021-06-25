
all: build

build:
	${MAKE} -C native
	npm install
	npm run package
	npm run make

.PHONY: all build