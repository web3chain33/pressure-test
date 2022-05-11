#!/bin/bash

# 获取编译参数
getFlags() {
  main_path="main"
  go_version=$(go version | awk '{ print $3 }')
  build_time=$(date "+%Y-%m-%d %H:%M:%S %Z")
  git_commit=$(git rev-parse --short=10 HEAD)
  builder_email=$(git config user.email)
  flags="-X '${main_path}.goVersion=${go_version}' -X '${main_path}.buildTime=${build_time}' -X '${main_path}.gitCommit=${git_commit}' -X '${main_path}.builderEmail=${builder_email}' -X 'google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn'"
  echo "${flags}"
}

# 创建目标目录
mkPkg() {
  file_path="$1"
  pkg_name="$2"

  mkdir -pv "${file_path}/${pkg_name}"
}

# 编译压测客户端
buildClient() {
  file_path="$1"
  pkg_name="$2"
  client_name="$3"

  flags=$(getFlags)

  cd "${file_path}"/chain33/"${client_name}" || exit
  echo "start building ${client_name} client"
  go build -ldflags "${flags}" -o "${file_path}/${pkg_name}/${client_name}" || exit
  echo "building ${client_name} client success"

  toml="${file_path}/chain33/${client_name}/config".toml
  if [ ! -f "${toml}" ]; then
    toml="${file_path}/chain33/${client_name}/config".toml
  fi
  cp "${toml}" "${file_path}/${pkg_name}/config".toml
}

# 打包目标目录
tarPkg() {
  file_path="$1"
  pkg_name="$2"
  file_name="$3"

  tar -zcvPf "${file_path}/${pkg_name}".tar.gz -C "${file_path}" "${pkg_name}"
  # rm -r "${file_path}/${pkg_name}"
}
