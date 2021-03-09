#!/bin/bash

export VERSION="0.1.2"

make generate
make manifests

export OPERATOR_IMG="quay.io/apahim/addon-operator:${VERSION}"
export BUNDLE_IMG="quay.io/apahim/addon-operator-bundle:${VERSION}"
export INDEX_IMG="quay.io/apahim/addon-operator-index:${VERSION}"

# publish the operator image
make docker-build docker-push IMG=$OPERATOR_IMG

# publish the bundle image
make bundle IMG=$OPERATOR_IMG
make bundle-build BUNDLE_IMG=$BUNDLE_IMG
make docker-push IMG=$BUNDLE_IMG

# publish the index image
opm index --container-tool docker add --bundles "${BUNDLE_IMG}" --tag "${INDEX_IMG}"
docker push "${INDEX_IMG}"
