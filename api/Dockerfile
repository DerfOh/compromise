FROM golang:latest
MAINTAINER Fredrick Paulin <fredrick.p@outlook.com>

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
EXPOSE 1234 1234
# Set production port
# EXPOSE 80:80

CMD ["/app/main"]
