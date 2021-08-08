# PEScanner

[![Build](https://github.com/petershen0307/PEScanner/actions/workflows/build.yml/badge.svg)](https://github.com/petershen0307/PEScanner/actions/workflows/build.yml)

This is a PE scanner and result with PE sha2 value.

single thread mode sample command
`go run main.go -mode=1 -entry=/mnt/c/Users/PC/Desktop/code/go/PEScanner -output=/mnt/c/Users/PC/Desktop/code/go/PEScanner/output`

concurrent mode sample command
`go run main.go -mode=2 -thread=3 -entry=/mnt/c/Users/PC/Desktop/code/go/PEScanner -output=/mnt/c/Users/PC/Desktop/code/go/PEScanner/output`
