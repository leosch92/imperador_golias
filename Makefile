GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w
GOGET=$(GOCMD) get
GOPATH=$(HOME)/go
GOBIN=$(GOPATH)/bin

#Packages

#MPKG := main


default: install

install:
	echo "Brew will be updated"
	brew update
	echo "Go will be installed"
	brew install golang
	echo "Go installed successfully"
	mkdir -p $(GOPATH)
	mkdir -p  $(GOBIN)
	echo "Pigeon will be installed"
	$(GOGET) -u github.com/mna/pigeon
	echo "Pigeon installed"
	#echo "Cloning project into local directory"
	#git clone https://github.com/leosch92/imperador_golias.git
	#echo "Project cloned successfully"
	$(GOBIN)/pigeon -o=src/parser.go peg/imp.peg
	#sh make.sh
	$(GOBUILD) -o=main ./src
	./main src/program.imp

main:
	$(GOBIN)/pigeon -o=src/parser.go peg/imp.peg
	$(GOBUILD) -o=main ./src
	./main src/program.imp

.PHONY: clean
clean:
	rm main
	rm src/parser.go

.PHONY: uninstall
uninstall:
	rm -r $(GOPATH)
	brew uninstall golang
	cd ../; rm -rf imperador_golias
