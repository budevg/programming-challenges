#!/bin/bash
set -x
stack exec doctest -- -isrc src/Course/$1.hs
