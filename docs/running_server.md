# How to run the server
Ensure enviroment variables are set
```bash
    go build main.go
    ./main
```

## An easier way
1. Create the config file. The contents should look like the following
```bash
    export DB_USER="FILL OUT THESE"       # user for mysql server
    export DB_PASSWD="FIELDS WITH INFO"   # password for said user
    export DB_LOC="THAT WILL NEED TO BE"  # probably localhost
    export DB_NAME="SPECIFIC PER MACHINE" # the name of the database to use

    export TOK_SECRET="LONG RANDOM STRING" # used for signing the tokens

    export DID_I_SET_THE_ENVIROMENT_VARIABLES="YES I DID" # to show an error message if you forget to set variables
```

2. Run the script
```bash
    ./run.sh
```

