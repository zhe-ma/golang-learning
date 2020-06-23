#!/bin/bash

# 1. #!告诉系统其后路径所指定的程序即是解释此脚本文件的 Shell 程序

# 2. chmod +x ./test.sh  #使脚本具有执行权限

# 3. echo
# echo "Hello World"

# 4. 调用语句
# echo $(ls /etc)
# echo `ls /etc`

# 5. 使用变量的时候要加$，但是给变量赋值时不能加，即使是第二次赋值。
# files=`ls /etc`
# echo $files
# echo ${files}

# 6. for循环
# for skill in CPP JAVA Go; do
#   echo "I am good at ${skill}!"
# done

# 7. 字符串
# 单引号里的任何字符都会原样输出，单引号字符串中的变量是无效的。
# 双引号里可以有变量，双引号里可以出现转义字符。

# name="AAA"
# greeting="Hello, "${name}"!" # 拼接字符串
# echo $greeting

# string="12345678"
# echo ${#string}  #  获取字符串长度
# echo ${string:1:4}  # 提取字符串第二到四个字符

# 8. 数组
# array=(1 2 3 4)
# array[0]=100
# echo ${array[0]}
# echo ${array[1]}
# echo ${#array[*]}  # 获取数组元素个数：4
# echo ${#array[0]}  # 获取单个元素的长度：3

# 9. 向脚本传递参数，脚本内获取参数的格式为：$n
# echo $0
# echo $1
# for i in "$*"; do  # $*将参数合并成一个
#   echo $i
# done

# for i in "$@"; do  # $@原来的参数个数
#     echo $i
# done

# 10. 基本运算符
# 原生bash不支持简单的数学运算，但是可以通过其他命令来实现。
# 例如 awk 和 expr，expr 最常用。
# expr 是一款表达式计算工具，使用它能完成表达式的求值操作。 
# a=10
# b=11
# c=a+b
# echo $c  # 输出a+b
# c=`expr $a + $b`  # 注意使用的是反引号 ` 而不是单引号 '
# echo $c  # 输出21
# val=`expr $a \* $b`  # 乘号(*)前边必须加反斜杠(\)才能实现乘法运算；
# echo $val

# 11. printf
# printf 命令模仿 C 程序库（library）里的 printf() 程序。
# printf "%-10s %-8s %-4s\n" 姓名 性别 体重kg  
# printf "%-10s %-8s %-4.2f\n" 郭靖 男 66.1234 
# printf "%-10s %-8s %-4.2f\n" 杨过 男 48.6543 
# printf "%-10s %-8s %-4.2f\n" 郭芙 女 47.9876 

# 12. If
# shell 语言中 0 代表 true，0 以外的值代表 false。
# a=10
# b=20
# if [ $a == $b ]  # 中括号必须有空格
# then
#   echo "a == b"
# elif [ $a -gt $b ]
# then
#   echo "a > b"
# elif [ $a -lt $b ]
# then
#   echo "a < b"
# else 
#   echo "None"
# fi

# 13. for
# array=("abc"  1234 "aaa")
# for i in ${array[*]}
# do
#   echo $i
# done

# for((i=1;i<=5;i++))
# do
#   echo $i
# done

# 14. while
# n=1
# while(( $n<=5 ))
# do
#   echo $n
#   let "n++"  #  Bash let 命令，它用于执行一个或多个表达式，变量计算中不需要加上 $ 来表示变量
# done

# echo 'CTRL-D退出'
# echo -n '输入'
# while read input
# do
#   echo "Hello ${input}"
# done

# 15. 函数
# func() {
#   echo "第一个参数$1"
#   echo "第二个参数$2"
#   echo "第三个参数$3"
# }

# func 100 102 103
