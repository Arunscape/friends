rm -f main
go build main.go

. .secret.sh
./main $1

rm -f main
