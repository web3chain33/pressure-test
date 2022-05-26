#!/bin/bash
file_path=$(
  cd "$(dirname "$0")" || exit
  pwd
)/..
# shellcheck source=./util.sh
source "${file_path}"/build/util.sh
client_name="$1"

project_name="pressure-${client_name}"
pkg_name="${project_name}"

mkPkg "${file_path}" "${pkg_name}"
buildClient "${file_path}" "${pkg_name}" "${client_name}"
tarPkg "${file_path}" "${pkg_name}" "${client_name}"
