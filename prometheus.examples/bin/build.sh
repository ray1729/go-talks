#!/bin/bash

for m in auto custom exercise-handler handler  random  red  simple vec; do
    cd $m
    go build -o ../bin/$m
    cd ..
done
