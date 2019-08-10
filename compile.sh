#!/bin/bash
file="$1"
g++ -std=c++11 "$1" -o out/${file/".cpp"/""}

