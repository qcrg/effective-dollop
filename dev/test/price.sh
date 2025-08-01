#!/bin/env bash
set -e

curl -kv -X GET "https://localhost:8643/currency/price?coin=bitcoin&timestamp=1722813300"
