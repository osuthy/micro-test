#!/bin/bash

set eu

docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=micro-test -p 3307:3306 mysql 
