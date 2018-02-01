#!/bin/bash


export GOPATH=`pwd`

echo 'Compiling and installing...'
go clean ui
go install ui
