#!/bin/zsh

# build reference implementation binary called "gemtext"

generate_tests() {
    go build -C golang .
    # generate html from reference implementation
    for f in tests/*.gmi; do
        golang/gemtext < "${f}" > "${f/.gmi/.html}"
    done
}

test_command() {
    # echo $1
    for f in tests/*.gmi; do
        # ${1}/gemtext < $f
        if ! diff -q <(${1}/gemtext < ${f}) <(cat ${f/.gmi/.html}); then
            echo ${1} $f
            exit 1
        fi
    done
}

generate_tests
test_command golang
test_command python
test_command elixir
