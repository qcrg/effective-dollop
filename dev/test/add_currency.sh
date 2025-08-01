#!/bin/env bash
set -e

curl -kvX POST "https://localhost:8643/currency/add" -H "Content-Type: application/json" -d '{"coin":"batcat"}'
