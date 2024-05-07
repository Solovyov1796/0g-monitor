#!/bin/bash

cd $(dirname ${BASH_SOURCE[0]})/..

pid=`ps -ef | grep "0g-monitor" | grep -v grep | awk '{print $2}'`
if [[ "$pid" != "" ]]; then
    echo "Terminate monitoring service, pid = $pid"
    kill $pid
fi

if ! [ -f .env ]; then
    echo "Error: .env file not found under `pwd`"
    exit 1
fi

source .env
./0g-monitor
