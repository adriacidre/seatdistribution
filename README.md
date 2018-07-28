### Getting the code

```
go get -u github.com/adriacidre/seatdistribution
cd $GOOPATH/src/github.com/adriacidre/seatdistribution
```

### Running the server

This software is provided with a Makefile, you should be able to perform basic tasks with it. Run `make` to see all options.

### Running tests / linters

Run package tests by executing `make test` on the root package folder.
To run linters you can do it with `make lint`

### Querying the API

Open 2 terminal consoles, on the first one run the server with `make serve`, on the other you can create a section and start assigning seats to it with:
```
# Adding a seat
$ make add-section
curl -X POST --data '{"id":"mysection","rows":1,"blocks":[3,4,3]}' localhost:9000/sections
{"success"}%                                          

# Assigning a new seat
$ make assign-seat
curl -X POST --data '{"id":"mysection"}' localhost:9000/sections/assign
"1C"%

# Getting a specific seat number
$ make number=1 get-seat-number
```

### LICENSE

[GNU AFFERO GENERAL PUBLIC LICENSE](LICENSE.md)
