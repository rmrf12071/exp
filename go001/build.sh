#!/bin/bash
# build on Linux(WSL2) for windows
GOOS=windows go build -ldflags="-s -w" -trimpath -o go001.exe .
