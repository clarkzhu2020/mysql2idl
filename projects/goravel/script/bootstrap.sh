#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=user
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}