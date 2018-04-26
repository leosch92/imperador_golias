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


default: main

main:
	echo "Brew will be updated"
	brew update
	echo "Go will be installed"
	brew install golang
	echo "Go installed successfully"
	mkdir -p $(HOME)/go
	mkdir -p  $(HOME)/go/bin
	echo "Pigeon will be installed"
	$(GOGET) -u github.com/mna/pigeon
	echo "Pigeon instalado"
	#echo "Cloning project into local directory"
	#git clone https://github.com/leosch92/imperador_golias.git
	#echo "Project cloned successfully"
	$(GOBIN)/pigeon -o=parser.go imp.peg
	#sh make.sh
	$(GOBUILD) main.go tree.go stack.go smc.go parser.go
	./main

.PHONY: clean
clean:
	rm -r imperador_golias
