#!/bin/bash
set -x
V=""
if [ "$1" == "-v" ]; then
    V="--quickcheck-verbose"
    shift
fi
stack test --test-arguments "--color=always -p test/Tasty.hs/$1 $V"
