echo "Iniciando processo..."
echo "Brew sera instalado. Caso ja possui Hombrew instalado, pressione qualquer tecla, exceto enter, para abortar. Caso contrario, pressione enter para continuar"

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
echo "Brew instalado com sucesso."

brew update
echo "Brew foi atualizado"

echo "Go sera instalado"
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

echo "Pigeon sera instalado"
go get -u github.com/mna/pigeon
echo "Pigeon instalado"

rm -rf imperador_golias

echo "Clonando projeto em diretoria local"
git clone https://github.com/leosch92/imperador_golias.git
echo "Projeto clonado"

cd imperador_golias

pigeon -o=parser.go imp.peg
echo "Parser gerado"

go build main.go tree.go stack.go smc.go parser.go
./main
