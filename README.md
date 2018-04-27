# Imperador_golias

<p align="center">
  <img src="imperador_golias.png">
</p>

Imperador Golias consegue a façanha de compilar E interpretar um código na linguagem imp (ou pelo menos no futuro irá)


Instruções (pt-BR):  
* Instalação - MacOS: para instalar Imperador Golias e todas as suas dependências, e rodar o projeto, é necessário ter o Homebrew instalado, então basta digitar `make` no terminal, na pasta onde está localizado o projeto.  
* Compilar e rodar - Linux/MacOS: já tendo todas as dependências (Go e Pigeon) instaladas, basta digitar `make {nome_do_arquivo}` no terminal, na pasta onde o projeto está localizado. Substituir {nome\_do\_arquivo} pelo nome do arquivo desejado, em /testes.  
    * fatorial.imp: fatorial (padrão: 5!)
    * pa\_soma\_n\_termos: soma dos n termos em uma progressão aritimética, iterativo (padrão: 10)
    * pa\_soma\_n\_termos2: soma dos n termos em uma progressão aritimética, por fórmula (padrão: 9)
* Limpar arquivos gerados - Linux/MacOS: digite `make clean` no terminal, na pasta onde o projeto se encontra.  
* Deinstalar projeto e todas as suas dependências - MacOS: desde que Imperador Golias tenha sido instalado de acordo com as instruções providenciadas nesse arquivo, basta usar `make uninstall` no terminal, na pasta em que o projeto se encontra.  


Instructions (en_US):  
* Installation - MacOS: to install Imperador Golias and all its dependencies, and run the project, it is necessary to have Homebrew installed. Then Imperador Golias can be ran by typing `make` in the terminal, in the folder where the project is located.  
* Compile and run - Linux/MacOS: if all dependencies are already installed (Go and Pigeon), all that is necessary is to type `make {filename}` in the terminal, in the folder where the project is located. Replace {filename} with the filename for the desired file, at /tests:  
    * fatorial.imp: factorial program (default: 5!)
    * pa\_soma\_n\_termos: sum of n terms in an arithmetic progression, by loop (default: 10)
    * pa\_soma\_n\_termos2: sum of n terms in an arithmetic progression, by formula (default: 9)
* Clean all generated files - Linux/MacOS: type `make clean` in the terminal, in the folder where the project is located.  
* Uninstall the project and all of its dependencies - MacOS: provided that Imperador Golias was installed according to the instructions provided in this file, all that is necessary is to use `make uninstall` in the terminal, in the folder where the project is located.  
