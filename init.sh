#!/bin/bash

mkdir -p input

touch input/$1.txt input/$1_test.txt

mkdir -p cmd/$1

cp templ.go cmd/$1/main.go
