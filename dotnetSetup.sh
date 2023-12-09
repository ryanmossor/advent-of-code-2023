#!/usr/bin/env bash

if [[ -z "$1" ]]; then
    echo "No project name provided"
    exit 1
fi

projName="$1"

dotnet new console --output "$1"

cat > "./$projName/Makefile" << EOF
build:
	dotnet build

clean:
	dotnet clean

restore:
	dotnet restore

watch:
	dotnet watch --project ./$projName.csproj run

start:
	dotnet run --project ./$projName.csproj

EOF
