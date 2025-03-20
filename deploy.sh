#! /bin/bash

if [ ! -f 'goleveldbadmin' ]; then
  echo 文件不存在! 待添加的安装包: 'goleveldbadmin'
  exit
fi

chmod +x goleveldbadmin

echo "goleveldbadmin..."
sleep 3
docker stop goleveldbadmin-goleveldbadmin-1

sleep 2
docker rm goleveldbadmin-goleveldbadmin-1

docker rmi goleveldbadmin-goleveldbadmin
echo ""

echo "goleveldbadmin packing..."
sleep 3
docker-compose -f docker-compose.yaml up -d
echo ""

echo "goleveldbadmin success..."