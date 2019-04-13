rm -f main
go build main.go

. .secret.sh
./main

rm -f main
