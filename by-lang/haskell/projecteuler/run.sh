#!/bin/bash

cabal sandbox init
cabal install primes split
cabal exec ghci $1
