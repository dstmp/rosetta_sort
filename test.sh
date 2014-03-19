#!/bin/sh

export GOPATH="$PWD/golang/"
go test daniellesucher.com/heapsort/
go test daniellesucher.com/quicksort/
python python/heapsort.py -v
python python/quicksort.py -v
