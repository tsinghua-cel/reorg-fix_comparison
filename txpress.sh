#!/bin/bash
docker stop txpress ; docker rm txpress
docker run -d -v ${PWD}/config/txpress-app.json:/root/app.json --name txpress  tscel/txpress:0627 --start
