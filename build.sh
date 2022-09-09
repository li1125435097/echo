#! /bin/bash

serverExist=`ls server | wc -l`

if [ $serverExist == "0" ]
then 
  go build src/server.go
else
  echo "项目已编译，跳过build"
fi