shell分割字符串：
1. string="hello, shell, haha"
2. array=(${string//,/ })  
for var in ${array[@]}
do
  echo $var
done

