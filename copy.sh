#!/bin/sh
if [ $# -lt 1 ]; then
    echo arg error: $*
    exit 1
fi

scp -p linenotify pi@$1:/home/pi/linenotify
