SUBDIRS := src #$(wildcard */.)

all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@

.PHONY: setup-build-env
setup-build-env:
	$(MAKE) -C src/native/gomodule setup-build-env

.PHONY: install-build-deps
install-build-deps:
	# nodegomodule actually depends on gomodule but that needs to be compiled first
	# $(MAKE) -C src/native/nodegomodule install-build-deps

.PHONY: all $(SUBDIRS)