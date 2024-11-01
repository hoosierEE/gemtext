#!/bin/zsh

# generate html from reference implementation (uncomment next line to regenerate tests):
# for f in tests/*.gmi; do go run g2h.go < $f > "${f/.gmi/.html}"; done

# compare with expected html
for f in tests/*.gmi; do
    if ! diff -q <(go run g2h.go < $f) "${f/.gmi/.html}"; then
        echo "exiting after first test failure"
        exit 1
    fi
done

echo "all tests passed!"
