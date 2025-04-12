#start with go base image
FROM golang:1.23.4-bullseye


WORKDIR /app

#copy your go's source code and front end files
COPY . /app

#install nessecary depedencies, build the go app
RUN go build -o server server.go

#expose ports for go server (e.g. 8080 port for backend)
EXPOSE 8080

#run the go server and run static files for the go app
CMD ["./server"]


LABEL maintainer="mgi"

LABEL version="1.0"