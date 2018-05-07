# Imperador_golias

<p align="center">
  <img src="imperador_golias.png">
</p>

Imperador Golias consegue a façanha de compilar E interpretar um código na linguagem imp (ou pelo menos no futuro irá)


Instruções (pt-BR):  
* Instalação - MacOS: para instalar Imperador Golias e todas as suas dependências, e rodar o projeto, é necessário ter o Homebrew instalado. Após confirmar que o Homebrew está presente, siga esses passos:
  * Do terminal, navegue até o diretório em que deseja instalar o projeto  
  * Digite `git clone https://github.com/leosch92/imperador_golias.git` e dê enter. Será criada uma nova pasta neste diretório, chamada "imperador_golias", contendo todos os arquivos do projeto.  
  * Digite `cd imperador_golias`, e dê enter.
  * Digite `make` e dê enter. Todas as dependências serão baixadas e instaladas, e ao final o programa fornecerá uma lista de arquivos \*.imp presentes, e solicitará que escolha um para testar. Digite o nome do programa, ou `exit` para sair.
* Compilar e rodar - Linux/MacOS: já tendo todas as dependências (Go e Pigeon) instaladas, basta digitar `make main` no terminal, na pasta onde o projeto está localizado. O programa fornecerá uma lista de programas *imp* presentes e solicitará que escolhe um deles para rodar. Digite o nome do programa e dê enter para rodar um programa, ou `exit` para sair. Para adicionar e testar um novo programa dessa forma, adicione seu arquivo ao diretório /tests, com a extensão \*.imp.
* Limpar arquivos gerados - Linux/MacOS: digite `make clean` no terminal, na pasta onde o projeto se encontra.  
* Deinstalar projeto e todas as suas dependências - MacOS: desde que Imperador Golias tenha sido instalado de acordo com as instruções providenciadas nesse arquivo, basta usar `make uninstall` no terminal, na pasta em que o projeto se encontra.  


Instructions (en-US):  
* Installation - MacOS: to install Imperador Golias and all its dependencies, and run the project, it is necessary to have Homebrew installed. After making sure Homebrew is present, you can install Imperador Golias by following those instructions:  
  * From the terminal browse to the directory where you want to install the project  
  * Type `git clone https://github.com/leosch92/imperador_golias.git` and hit enter. A new folder will be created inside this directory, with the name "imperador_golias", containing all the project files.  
  * Type `cd imperador_golias` and hit enter  
  * Type `make` and hit enter. All dependencies will be downloaded and installed, and at the end the program will dispaly a list of \*.imp files present and will prompt you to type the name of the program you wish to run. Type the name of the program, or `exit` to exit.
* Compile and run - Linux/MacOS: if all dependencies are already installed (Go and Pigeon), all that is necessary is to type `make main` in the terminal, in the folder where the project is located. The program will dispaly a list of \*.imp files present and will prompt you to type the name of the program you wish to run. Type the name of the program, or `exit` to exit.
* Clean all generated files - Linux/MacOS: type `make clean` in the terminal, in the folder where the project is located.  
* Uninstall the project and all of its dependencies - MacOS: provided that Imperador Golias was installed according to the instructions provided in this file, all that is necessary is to use `make uninstall` in the terminal, in the folder where the project is located.  
