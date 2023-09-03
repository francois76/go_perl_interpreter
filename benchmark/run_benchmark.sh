#!/bin/bash
p=$(dirname "$0")
path=$(realpath "$p")
GREEN="\e[32m"
BOLD="\e[1m"
ENDCOLOR="\e[0m"

(cd "$path" && go build 1_full_go.go)


echo -e $(printf "${BOLD}${GREEN}Full go${ENDCOLOR}")
time "$path/1_full_go"

echo -e $(printf "${BOLD}${GREEN}Full perl${ENDCOLOR}")
time perl "$path/2_full_perl.pl"