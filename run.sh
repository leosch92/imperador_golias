dir=tests
cd $dir
dirlist=( `ls -B *.imp | sed -e 's/\..*$//'` )
#fcount='ls -B *.imp | wc - 1'
echo "Programas encontrados:"
for file in ${dirlist[*]}
do
  echo $file
done
cd ../
echo 'Qual programa deseja executar?'
read op
for file in ${dirlist[*]}
do
  if [ "$file" == "$op" ]
  then
    ./main "tests/$file.imp" -verbose
  fi
done
