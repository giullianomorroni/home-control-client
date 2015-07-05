#!/bin/bash
while true
do
	go run /opt/goapp/src/home-control-client/main.go >> /opt/goapp/src/home-control-client/home-control.log
	sleep 3
done
