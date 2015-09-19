# Shutils

**shutils** is a set of small shell utilities written in Go.

## Install

You need to have Go in order to compile the files.

    git clone https://github.com/bfontaine/shutils.git
    cd shutils && make install

## What’s in the box

* `sum` reads numbers on `stdin` and prints the sum of all of them. It ignores
  non-numbers and support floats with the `-F` option.
