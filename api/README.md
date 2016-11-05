#REST API in Golang with mySql Database

# Install go lang

# Installation

## Set up mysql Database
        create database compromise;
        use use compromise;

        * Then insert the tables located in the databse directory
        * After that run the test data

## Clone repo
        git clone https://github.com/derfoh/compromise
        cd compromise/golang/
        go run server.go

And open http://IP_or_localhost:1234/api

# User API Spec

GET /api/ to get all the Users.

POST /api/ to add new Users {EmailAddress,  FirstName, LastName, Nickname, Password}

DELETE /api/EmailAddress to remove that one User.

PUT /api/ to update details {EmailAddress,  FirstName, LastName, Nickname, Password}
