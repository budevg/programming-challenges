#!/bin/bash

stack exec \
      -- doctest -isrc \
      -Wall -fno-warn-unused-binds -fno-warn-unused-do-bind \
      -fno-warn-unused-imports -fno-warn-type-defaults \
      -ferror-spans -fno-warn-type-defaults \
      src/Course/$1.hs
