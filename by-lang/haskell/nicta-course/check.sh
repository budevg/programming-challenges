#!/bin/bash

stack exec -- doctest -isrc -Wall -fno-warn-type-defaults src/Course/$1.hs
