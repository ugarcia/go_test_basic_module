#!/usr/bin/env bash
SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"
HOME_DIR="${SCRIPT_DIR}/.."
docker run "$@" -d -p 58201:8201 -v ${HOME_DIR}:/go/src/github.com/ugarcia/go_test_basic_module go_test_basic_module
