#!/bin/zsh

# build reference implementation binary called "gemtext"
go build .

# generate html from reference implementation
for f in tests/*.gmi; do
    ./gemtext < $f > "${f/.gmi/.html}"
done

# compare with expected html
for f in tests/*.gmi; do
    if ! diff -q <(./gemtext < $f) "${f/.gmi/.html}"; then
        echo "exiting after first test failure"
        exit 1
    fi
done

echo "all tests passed!"
