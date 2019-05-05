#!/bin/bash
AUTH_SERVER="auth-server"
MSG_SERVER="msg-server"
EMAIL_SERVER="email-server"

function docker-compose-restart () {
  echo "Restarting service $1"
  docker-compose build $1
  docker-compose stop $1
  docker-compose up -d --no-deps $1
}

function deploy () {
  docker-compose-restart $AUTH_SERVER-$1
  docker-compose-restart $MSG_SERVER-$1
  docker-compose-restart $EMAIL_SERVER-$1
}

function full-redeploy() {
    docker-compose down
    racket make-docker-compose.rkt > docker-compose.yml
    docker-compose build
    docker-compose up -d
}

./build.sh $1
case "$1" in
  "")
    full-redeploy
    ;;
  "prod")
    deploy "prod"
    ;;
  "dev")
    deploy "dev"
    ;;
  "spike")
    deploy "spike"
    ;;
esac


