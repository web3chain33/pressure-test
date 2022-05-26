#!/bin/bash
file_path=$(
  cd "$(dirname "$0")" || exit
  pwd
)/..

rm -rf "${file_path}"/pressure-*
