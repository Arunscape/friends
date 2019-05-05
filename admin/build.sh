#!/bin/bash

LOC="/home/jacob/friends"

AUTH_SERVER="auth_server"
MSG_SERVER="msg_server"
EMAIL_SERVER="email_server"

PROD_BRANCH="server"
DEV_BRANCH="server"
SPIKE_BRANCH="server"

function build_all () {
  cd $LOC
  git checkout $2
  docker_build $1 $AUTH_SERVER &
  docker_build $1 $MSG_SERVER &
  docker_build $1 $EMAIL_SERVER &
}
function docker_build() {
  docker build --build-arg SERVER=$2 -f $LOC/admin/Dockerfile.server -t $1$2 . --rm
}

case "$1" in
  "")
    build_all "prod_" $PROD_BRANCH
    build_all "dev_" $DEV_BRANCH
    build_all "spike_" $SPIKE_BRANCH
    ;;
  "prod")
    build_all "prod_" $PROD_BRANCH
    ;;
  "dev")
    build_all "dev_" $DEV_BRANCH
    ;;
  "spike")
    build_all "spike_" $SPIKE_BRANCH
    ;;
esac

wait
echo "Y" | docker system prune --remove-orphans
