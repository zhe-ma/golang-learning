#!/bin/bash

# dir=`eval pwd`
dir="/mnt/c/Users/Admin/AppData/Local/Packages/Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy/LocalState/Assets"
dest="/mnt/d/wallpapers"

echo "Copy images from $dir."

for file in $(ls $dir)
do
  size=`file $dir/$file | cut -d ',' -f 8`
  echo "size:$size"

  width=${size:1:4}
  height=${size:6:9}

  #需要转义<，否则认为是一个重定向符号
  if [ $width \> $height ];
  then
    new="$dest/${file:0:10}.png"
    cp "$dir/$file" $new 
    echo "Copied: $new"
  fi
done