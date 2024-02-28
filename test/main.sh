#!/bin/bash

echo 'Running test:' 

echo 'JavaScript:'
node main.js

sleep 2

echo 'Go:'

go run main.go
