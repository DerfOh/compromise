FROM golang:latest
MAINTAINER Fredrick Paulin <fredrick.p@outlook.com>

# TODO: add ENV variables to parameterize run command 
# /app/main -dbaddress=198.199.75.80 -dbuser=compromise -dbpassword=password
ENV DBADDRESS localhost
ENV DBUSER compromise
ENV DBPASSWORD password
# ENV APIPORT 80
ENV APIPORT 443

# Create the directory for the application to run from
RUN mkdir /app 

# Add the contents of the current directory to the dockerimage
ADD . /app/ 

#Set your working directory
WORKDIR /app 

# Retrieve the files for mysql from github
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt
RUN go get golang.org/x/crypto/acme/autocert

# Build the application
RUN go build -o main . 

# Set Test port
# EXPOSE 1234 1234

# Set production port, SSL is handled on loadbalancer
EXPOSE 80 80
#EXPOSE 443 443

CMD ["sh", "-c","/app/main -port ${APIPORT} -dbaddress ${DBADDRESS} -dbuser ${DBUSER} -dbpassword ${DBPASSWORD} Docker "]

# Example command for testing the container locally
# docker run -it -e DBADDRESS='127.0.0.1' -e DBPASSWORD='password' -e port=1234 -p 80:80 derfoh/compromise