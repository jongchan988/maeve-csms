#!/bin/sh


./app serve > server.log 2>&1 &
sleep 1
tail -f server.log