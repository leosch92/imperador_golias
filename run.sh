<<<<<<< HEAD
dir=tests
cd $dir
dirlist=( `ls -B *.imp | sed -e 's/\..*$//'` )
=======
#!/bin/bash
dir=tests
cd $dir
declare -a dirlist=( `ls -B *.imp | sed -e 's/\..*$//'` )
>>>>>>> enhancement/peg_com_clauses
#fcount='ls -B *.imp | wc - 1'
echo "Programas encontrados:"
for file in ${dirlist[*]}
do
  echo $file
done
cd ../
echo 'Qual programa deseja executar?'
read op
<<<<<<< HEAD
if [ "$op" == "exit" ]
then
    exit 0;
fi
=======
>>>>>>> enhancement/peg_com_clauses
for file in ${dirlist[*]}
do
  if [ "$file" == "$op" ]
  then
    ./main "tests/$file.imp" -verbose
  fi
done
