#!/bin/sh

cancel(){
    GOTODO_PID=$( cat .gotodo.pid )
    kill -9 $GOTODO_PID
}

run(){
    dep ensure
    go run main.go &
    GOTODO_PID=$( ps -fea | grep "[m]ain.go" | cut -d" " -f2 )
    echo $GOTODO_PID > .gotodo.pid
}

watch(){
    while inotifywait -r -e modify . ; do
    cancel
    run
    done
}

run
watch