#!/usr/bin/env bash
set -euo pipefail

LOGFILE=/tmp/mylog.log


function parseLog {
    ST=$1
    EN=$2
    LF=$3

    cat $LF | awk -v ST=$ST -v EN=$EN -F ':' '
{
if ($1 ~ ST) {st=$2}
if ($1 ~ EN) {en=$2; count=count+1; sum=sum+en-st}
}
END {print "From " ST " to " EN " ->Average value : " sum/count}
'
}
parseLog "START" "STOP" $LOGFILE
echo "---Details---"
parseLog "START" "INIT" $LOGFILE
parseLog "INIT" "PARSED" $LOGFILE
parseLog "PARSED" "BUILDED" $LOGFILE
parseLog "BUILDED" "ENDED" $LOGFILE
parseLog "ENDED" "STOP" $LOGFILE
