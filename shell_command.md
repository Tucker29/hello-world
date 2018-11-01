shell分割字符串：
1. string="hello, shell, haha"
2. array=(${string//,/ })  
for var in ${array[@]}  
do  
  echo $var  
done  
如果原来的字符串有空格： eg: ’mark:x:0:0:this is a test user:/var/mark:nologin’  
要将:分割的字符串输出, 上面的方法会将this is a test user分别输出,这个是不对的。
solution:

user=’mark:x:0:0:this is a test user:/var/mark:nologin’  
for((i=7;i<=7;i++))  
do  
  echo $user|cut -d ":" -f$i
done  
上述方法是确定了字符串分割后的长度的，要写的更通用一点的话：  
#!/bin/bash
user=’mark:x:0:0:this is a test user:/var/mark:nologin’
i=1
while((1==1))
do
        split=`echo $user|cut -d ":" -f$i`
        if [ "$split" != "" ]
        then
                ((i++))
                echo $split
        else
                break
        fi
done


