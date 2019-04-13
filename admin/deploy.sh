AUTH_SERVER="auth-server"
MSG_SERVER="msg-server"

function docker-compose-restart () {
  echo "Restarting service $1"
  docker-compose build $1
  docker-compose stop $1
  docker-compose up -d --no-deps $1
}

function deploy () {
  docker-compose-restart $AUTH_SERVER-$1
  docker-compose-restart $MSG_SERVER-$1
}

./build.sh $1
case "$1" in
  "")
    deploy "prod"
    deploy "dev"
    deploy "spike"
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


