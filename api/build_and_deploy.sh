#!/bin/bash

# Build go binary, name it api and move it to home directory
go build -o ~/api

# copy the index page to the the home directory for the api
cp index.html ~/

# start the app using port 8080 (production)
exec ~/api -port 8080

# start the app using port 1234 (test)
# exec ~/api -port 1234
