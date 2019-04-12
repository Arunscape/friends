LOC="/home/jacob/friends"

API_SERVER="api_server"

function build_all () {
  cd $LOC
  git checkout $2
  docker build -f $LOC/admin/Dockerfile.api_server -t $1$API_SERVER . --rm
}
case "$1" in
  "")
    build_all "prod_" devops
    build_all "dev_" devops
    build_all "spike_" devops
    ;;
  "prod")
    build_all "prod_" devops
    ;;
  "dev")
    build_all "dev_" devops
    ;;
  "spike")
    build_all "spike_" devops
    ;;
esac

echo "Y" | docker system prune
