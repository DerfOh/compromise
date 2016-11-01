#REST API in Golang with mySql Database

# Install go lang

# Installation

## Set up mysql Database
        create database farm;
        use farm;
        create table pandas (ID int NOT NULL AUTO_INCREMENT, name varchar(20), primary key (ID));
        
## Clone repo
        git clone https://github.com/motyar/restgomysql
        cd restgomysql
        go run server.go

And open http://IP_or_localhost:1234/api

# Nothing but (cute) Pandas

GET /api/ to get all the pandas.

POST /api/ to add new panda {name}

DELETE /api/panda_id to remove that one panda.

PUT /api/ to update details {id and name}



