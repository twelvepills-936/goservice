#!/usr/bin/env bash

set -ex

PACKAGE_NAME=$(go list -m)
PACKAGE_DIR=${GOPATH}/src/${PACKAGE_NAME}

# out generator dirs
DST_DIR_GO=${PACKAGE_DIR}/pkg/api
DST_DIR_PHP=${PACKAGE_DIR}/pkg/php_client

# flags
GEN_PHP=false
GEN_GO=false

# proto import paths
INCLUDE="-I${PACKAGE_DIR}/api -I. -I/usr/local/include -I${GOPATH}/src"

# import all proto dependencies
# generate proto files
# generate openapi file
function gen() {
  if [[ ${GEN_PHP} == false && ${GEN_GO} == false ]]; then
    GEN_GO=true
  fi

  # Создаем директории если их нет
  mkdir -p ${DST_DIR_GO}
  mkdir -p ${DST_DIR_PHP}
  mkdir -p ${PACKAGE_DIR}/api

  if [[ ${GEN_GO} == true ]]; then
    echo "Generating Go code..."
    rm -rf ${DST_DIR_GO}/*
    protoc_gen_code_go
  fi

  if [[ ${GEN_PHP} == true ]]; then
    echo "Generating PHP code..."
    rm -rf ${DST_DIR_PHP}/*
    protoc_gen_code_php
  fi

  echo "Generating OpenAPI spec..."
  protoc_gen_openapi
}

function init() {
  flags "$@"
  import_proto_deps
}

# generate code
function protoc_gen_code_go() {
  protoc ${INCLUDE} \
    --go_out=${DST_DIR_GO} \
    --go_opt=paths=source_relative \
    --go_opt=default_api_level=API_OPAQUE \
    --go-grpc_out=${DST_DIR_GO} \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=${DST_DIR_GO} \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --validate_out=lang=go:${DST_DIR_GO} \
    --validate_opt=paths=source_relative \
    ${PACKAGE_DIR}/api/*.proto  # ✅ Генерируем из всех proto файлов
}

# generate code PHP
function protoc_gen_code_php() {
  protoc ${INCLUDE} \
    --php_out=${DST_DIR_PHP} \
    ${PACKAGE_DIR}/api/*.proto  # ✅ Генерируем из всех proto файлов

  #Удаление подключения классов с остутвием автогенерации
  find ${DST_DIR_PHP} -type f -exec sed -i -e '/Envoyproxy\\ProtocGenValidate/d' {} \;
}

# generate service.swagger.json
function protoc_gen_openapi() {
  # Генерируем OpenAPI для всех proto файлов или только для основных
  protoc ${INCLUDE} \
    --openapiv2_out=${PACKAGE_DIR}/api \
    ${PACKAGE_DIR}/api/*.proto  # ✅ Или укажите конкретные файлы
}

function import_proto_deps {
  go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/envoyproxy/protoc-gen-validate

  go get github.com/googleapis/googleapis

  proto_deps=(
    "google.golang.org/protobuf"
    "github.com/googleapis/googleapis"
    "github.com/grpc-ecosystem/grpc-gateway/v2"
    "github.com/envoyproxy/protoc-gen-validate"
  )

  for dep in "${proto_deps[@]}"; do
    pkg_name=${dep%>*}
    pkg_proto_dir=${dep#*>}

    pkg_path="$(go list -m -f '{{ .Dir }}' -mod=mod "${pkg_name}")"
    if [[ "${pkg_name}" != "${pkg_proto_dir}" ]]; then
      pkg_path="${pkg_path}/${pkg_proto_dir}"
    fi

    INCLUDE="${INCLUDE} -I${pkg_path}"
  done

  PATH="/go/bin":${PATH}
  export PATH
}

function flags() {
  for arg in "$@"; do
    case ${arg} in
    --php)
      GEN_PHP=true
      shift
      ;;
    --go)
      GEN_GO=true
      shift
      ;;
    esac
  done
}

init "$@"
gen
