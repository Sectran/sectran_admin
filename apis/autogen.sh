#!/bin/bash

username="root"
password="Ryan@1218pass"
database="sectran"
target_dir="../apiservice"

goctl api go --api="apis.api" --style="goZero" --dir=$target_dir
tables=$(mysql -uroot -pRyan@1218pass -e "use sectran;show tables" | grep st_ | grep -v grep)

for table in ${tables}; do
  echo "Generating model for table: $table"
  goctl model mysql datasource -dir="$target_dir/model/$table" -url="$username:$password@tcp(127.0.0.1:3306)/$database" -style=goZero -table="$table"
done