#!/bin/bash
currDir=$(dirname $0)
(cd "$currDir"/"$1" && go run $1.go)