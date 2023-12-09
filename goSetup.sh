#! /usr/bin/env bash

if [[ -z "$1" ]]; then
    echo "No project name provided"
    exit 1
fi

projName="$1"

mkdir -p "$1"
cp ./utils.go "$1"

cd "$1"

touch "$1".go

