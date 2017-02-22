FROM golang:latest
MAINTAINER Fredrick Paulin <fredrick.p@outlook.com>

# TODO: add ENV variables to parameterize run command 
# /app/main -dbaddress=198.199.75.80 -dbuser=compromise -dbpassword=password
ENV DBADDRESS localhost
ENV DBUSER compromise
ENV DBPASSWORD password
ENV APIPORT 80

# Create the directory for the application to run from
RUN mkdir /app 

# Add the contents of the current directory to the dockerimage
ADD . /app/ 

#Set your working directory
WORKDIR /app 

# Retrieve the files for mysql from github
RUN go get github.com/go-sql-driver/mysql

# Build the application
RUN go build -o main . 

# Set Test port
# EXPOSE 1234 1234
# Set production port
EXPOSE 80 80

CMD ["sh", "-c","/app/main -port ${APIPORT} -dbaddress ${DBADDRESS} -dbuser ${DBUSER} -dbpassword ${DBPASSWORD} Docker "]
