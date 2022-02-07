#!/usr/bin/env bash

set -e

echo "--- yarn in root"
# mutex is necessary since CI runs various yarn installs in parallel
pnpm --frozen-lockfile

echo "--- generate"
pnpm gulp generate

cd "$1"
echo "--- test"

# Limit the number of workers to prevent the default of 1 worker per core from
# causing OOM on the buildkite nodes that have 96 CPUs. 4 matches the CPU limits
# in infrastructure/kubernetes/ci/buildkite/buildkite-agent/buildkite-agent.Deployment.yaml
pnpm -s run test --maxWorkers 4 --verbose
