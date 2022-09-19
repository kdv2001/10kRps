FROM golang:latest

WORKDIR /server
COPY . .
RUN apt-get update
RUN go mod download
RUN go build -o server ./cmd
EXPOSE 8010
ENTRYPOINT ./server