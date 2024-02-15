#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=grom
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}