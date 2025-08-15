#!/bin/zsh

# build reference implementation binary called "gemtext_ref"

generate_tests() {
    # generate html from reference implementation
    go build -C golang .
    for f in tests/*.gmi; do
        golang/gemtext_ref < "${f}" > "${f/.gmi/.html}"
    done
}

test_command() {
    for f in tests/*.gmi; do
        if ! diff -q <(${1}/gemtext < ${f}) ${f/.gmi/.html}; then
            echo ${1} $f
            exit 1
        fi
    done
}

generate_tests
test_command python
test_command elixir
