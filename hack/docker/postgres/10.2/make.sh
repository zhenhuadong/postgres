#!/bin/bash
set -xeou pipefail

GOPATH=$(go env GOPATH)
REPO_ROOT=$GOPATH/src/github.com/kubedb/postgres

source "$REPO_ROOT/hack/libbuild/common/lib.sh"
source "$REPO_ROOT/hack/libbuild/common/kubedb_image.sh"

IMG=postgres
TAG=10.2

WALG_VER=${WALG_VER:-v0.1.7}

DIST="$REPO_ROOT/dist"
mkdir -p "$DIST"

build_binary() {
  pushd $REPO_ROOT
  ./hack/builddeps.sh
  ./hack/make.py build pg-operator
  popd
}

build_docker() {
  pushd "$REPO_ROOT/hack/docker/postgres/$TAG"

  # Download wal-g
  wget https://github.com/kubedb/wal-g/releases/download/${WALG_VER}/wal-g-alpine-amd64
  chmod +x wal-g-alpine-amd64
  mv wal-g-alpine-amd64 wal-g

  # Copy pg-operator
  cp "$DIST/pg-operator/pg-operator-alpine-amd64" pg-operator
  chmod 755 pg-operator

  local cmd="docker build -t $DOCKER_REGISTRY/$IMG:$TAG ."
  echo $cmd; $cmd

  rm wal-g pg-operator
  popd
}

build() {
  build_binary
  build_docker
}

binary_repo $@
