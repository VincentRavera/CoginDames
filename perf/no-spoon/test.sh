#!/usr/bin/env bash
set -euo pipefail

echo "# PERF START: $(date '+%s%N' | cut -b1-13)"
cat ./input.txt | clojure ../../me_there_is_no_spoon.clj 2>&1 | grep '# PERF'
echo "# PERF STOP: $(date '+%s%N' | cut -b1-13)"
