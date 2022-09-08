#! /bin/bash

# golang版本检测
goMod=`go env GO111MODULE`
expectMod="on"
if [ $goMod != $expectMod ]
then  
  echo "项目mod开启检测异常，程序退出"
  echo "请确保go版本在1.11以上，并打开gomod模式"
  echo "gomod开启命令为： go env -w GO111MODULE=on"
  exit 0
else
  echo "项目mod开启检测：无异常"
fi





# 项目位置检测
goPath=`go env GOPATH`
expectWorkPlace="$goPath/src"
currentDir=`cd .. && pwd`

if [ $expectWorkPlace != $currentDir ]
then 
  echo "项目位置检测异常，程序退出"
  echo "你的项目存放位置不正确，请把项目放到 $expectWorkPlace"
  echo "操作命令为： cd .. && mv echo $expectWorkPlace && cd $expectWorkPlace/echo"
  exit 0
else
  echo "项目位置检测：无异常"
fi


go run src/server.go