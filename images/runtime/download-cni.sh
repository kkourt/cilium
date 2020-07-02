#!/bin/bash

# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

set -o xtrace
set -o errexit
set -o pipefail
set -o nounset

cni_version="0.7.5"

for arch in amd64 arm64 ; do
  curl --fail --show-error --silent --location "https://github.com/containernetworking/plugins/releases/download/v${cni_version}/cni-plugins-${arch}-v${cni_version}.tgz" --output "/tmp/cni-${arch}.tgz"
  mkdir -p "/out/linux/${arch}/bin"
  tar -C "/out/linux/${arch}/bin" -xf "/tmp/cni-${arch}.tgz" ./loopback
done

strip /out/linux/amd64/bin/loopback
aarch64-linux-gnu-strip /out/linux/arm64/bin/loopback
