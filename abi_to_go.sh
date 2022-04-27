#!/bin/bash

echo "ABI to Go File Generator"
echo -n "Enter path to Contract ABI/JSON: "
read file_path
echo -n "Enter Contract Name: "
read abi_name
echo "⏳ Loading $abi_name ..."

mkdir -p ./abis/$abi_name
abigen --pkg main --abi $file_path --out ./abis/$abi_name/$abi_name.go

if [ $? -eq 0 ]; then
    echo "✅ $abi_name loaded successfully"
else
    rm -rf ./abis/$abi_name
    echo "❌ $abi_name failed to load"
fi