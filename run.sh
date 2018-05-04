dir=tests
cd $dir
dirlist=(`ls -B *.imp | sed -e 's/\..*$//'`)
#fcount='ls -B *.imp | wc - 1'
for file in ${dirlist[*]}
do
  echo $file
done
cd ../
read -p "Qual programa deseja executar?" op
for file in ${dirlist[*]}
do
  if [ "$file" == "$op" ]
  then
    ./main "tests/$file.imp" -verbose
  fi
done
