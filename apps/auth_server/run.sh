trap 'rm -f main' INT

go build main.go

. .secret.sh
./main $1

rm -f main
