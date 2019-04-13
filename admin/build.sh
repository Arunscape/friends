LOC="/home/jacob/friends"

AUTH_SERVER="auth_server"

PROD_BRANCH="server"
DEV_BRANCH="server"
SPIKE_BRANCH="server"

function build_all () {
  cd $LOC
  git checkout $2
  docker build -f $LOC/admin/Dockerfile.auth_server -t $1$AUTH_SERVER . --rm
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

echo "Y" | docker system prune
