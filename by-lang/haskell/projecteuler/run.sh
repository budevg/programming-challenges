#!/bin/bash

cabal sandbox init
cabal install primes
cabal exec ghci $1
