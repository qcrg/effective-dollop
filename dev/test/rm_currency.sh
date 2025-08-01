#!/bin/env bash
set -e

curl -kv \
  -X POST \
  "https://localhost:8643/currency/remove" \
  -H "Content-Type: application/json" \
  -d '{"coin":"BitCoiN"}'
