#!/bin/bash

CGO_ENABLED=0 go build -ldflags "-s -w" main.go 
echo "build completed"
mv main flightcheck-$1
echo "transfering..."
scp flightcheck-$1 root@62.106.95.203:
