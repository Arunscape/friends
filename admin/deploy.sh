API_SERVER="api-server"

function docker-compose-restart () {
  echo "Restarting service $1"
  docker-compose build $1
  docker-compose stop $1
  docker-compose up -d --no-deps $1
}

function deploy () {
  ./build.sh $1
  SUFFIX=""
  if [ "dev" == "$1" ]; then
    SUFFIX="1"
  fi
  if [ "spike" == "$1" ]; then
    SUFFIX="2"
  fi

  docker-compose-restart $API_SERVER$SUFFIX
}

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


