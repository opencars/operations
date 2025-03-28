.PHONY: default all clean
APPS     := server worker
BLDDIR   := bin
VERSION  := $(shell cat VERSION)
IMPORT_BASE := github.com/opencars/operations
LDFLAGS  := -ldflags "-X $(IMPORT_BASE)/pkg/version.Version=$(VERSION)"

default: clean all

all: $(APPS)

$(BLDDIR)/%:
	go build $(LDFLAGS) -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done
