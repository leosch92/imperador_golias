echo "Iniciando processo..."

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
echo "Brew instalado com sucesso."

brew update
echo "Brew foi atualizado"

brew install golang

echo "Go instalado com sucesso"

export PATH=/usr/local/go/bin:$PATH
mkdir $HOME/go
mkdir $HOME/go/bin
export GOPATH=$HOME/go
echo $GOPATH
export GOBIN=$GOPATH/bin
echo $GOBIN
export PATH=$GOBIN:$PATH
echo $PATH
echo "Env vars criadas"

go get -u github.com/mna/pigeon
echo "Pigeon instalado"

git clone https://github.com/leosch92/imperador_golias.git
echo "Projeto clonado"

cd imperador_golias

pigeon -o=parser.go imp.peg
echo "Parser gerado"

go build main.go tree.go stack.go smc.go parser.go
./main
