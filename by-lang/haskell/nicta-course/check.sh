#!/bin/bash

stack exec -- doctest -isrc -Wall \
      -fno-warn-unused-top-binds \
      -fno-warn-orphans \
      -fno-warn-type-defaults \
      -fno-warn-unused-imports \
      -fno-warn-redundant-constraints \
      src/Course/$1.hs
