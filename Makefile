.PHONY: default all clean
APPS     := http-server grpc-server worker
BLDDIR   := bin
VERSION  := $(shell cat VERSION)
IMPORT_BASE := github.com/opencars/operations
LDFLAGS  := -ldflags "-X $(IMPORT_BASE)/pkg/version.Version=$(VERSION)"

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: clean all

all: $(APPS)

$(BLDDIR)/%:
	go build --race $(LDFLAGS) -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done
