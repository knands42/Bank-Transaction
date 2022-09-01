FROM golang:1.19.0-bullseye

RUN apt update && apt upgrade -y
RUN apt install bash curl make gcc g++ -y

ENV PATH="$PATH:/bin/bash"

WORKDIR /go/Database-Transactions-Simulation