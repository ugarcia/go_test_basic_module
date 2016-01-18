#!/usr/bin/env bash
SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"
HOME_DIR="${SCRIPT_DIR}/.."
docker run "$@" -d -p 58201:8201 -v ${HOME_DIR}:/go/src/app go_test_basic_module
