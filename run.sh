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
    while inotifywait -r --exclude "(\.git|run.sh|vendor|Dockerfile|docker-compose.yml|\.dockerignore|debug|Gopkg.lock|README.md|\.gotodo\.pid)" -e modify . ; do
    cancel
    run
    done
}

/usr/bin/wait-for-it db:27017 -t 0 -- echo "MongoDB Started"
run
watch