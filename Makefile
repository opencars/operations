.PHONY: default all clean
APPS     := server worker
BLDDIR   := bin
VERSION  := $(shell cat VERSION)
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
