#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}
echo $CODEGEN_PKG

${CODEGEN_PKG}/generate-groups.sh all \
	github.com/interma/programming-k8s/pkg/client \
	github.com/interma/programming-k8s/pkg/apis \
  	"stats:v1alpha1" \
  	--go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt


