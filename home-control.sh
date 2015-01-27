#!/bin/bash

while true
do
	go run main.go >> home-control.log
	sleep 5
done
