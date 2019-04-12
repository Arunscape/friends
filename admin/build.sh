LOC="/home/jacob/friends"

API_SERVER="api_server"

function build_all () {
  cd $LOC
  git checkout $2
  docker build -f $LOC/admin/Dockerfile.api_server -t $1$API_SERVER . --rm
}

build_all "prod_" master
build_all "dev_" master
build_all "spike_" devops

echo "Y" | docker system prune
