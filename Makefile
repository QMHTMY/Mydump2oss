#Go related variables
GOCMD=go
GOVET=$(GOCMD) vet
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build

#Project related variables
BIN="bin"
SRC=$(shell find . -name "*.go")
OBJECT="Mydump2oss"
INSTALLDIR="/usr/bin"
GENMKDPAGE="docs/mkdPage.go"

#Make work flow
default: all

all: vet test fmt bin build docs

vet:
	$(info ******************** vetting ********************)
	$(GOVET) ./...

test:
	$(info ******************** testting ********************)
	$(GOTEST) -v ./...

fmt:
	$(info ******************** formatting ********************)
	gofmt -w $(SRC)

bin: 
	$(info ******************** make dir ********************)
	mkdir $(BIN)

build:
	$(info ******************** building ********************)
	$(GOBUILD) -o $(BIN)/$(OBJECT)

docs:
	$(info ******************** generating markdown page ********************)
	$(GORUN) $(GENMKDPAGE) 

install:
	install -d $(INSTALLDIR)
	sudo install -m 0755  $(BIN)/$(OBJECT) $(INSTALLDIR)

uninstall:
	sudo rm $(INSTALLDIR)/$(OBJECT)

clean:
	rm -rf $(BIN)

.PHONY: bin fmt vet test build docs install uninstall clean
