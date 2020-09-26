#!/usr/bin/env bash
set -euo pipefail

rm /tmp/mylog.log
touch /tmp/mylog.log
for i in $(seq 50)
do
    echo -n -e "$i\r"
    bash test.sh >> /tmp/mylog.log
done
echo
echo "DONE"
