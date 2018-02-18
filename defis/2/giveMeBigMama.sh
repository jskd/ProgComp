#!/usr/bin/env bash

ocamloptp -o bigmama $(dirname $0)/generateBigMama.ml
./bigmama 100 50000

