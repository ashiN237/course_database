FROM golang:1.18

RUN apt-get update && \
    apt-get install -y \
        git curl vim

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /workspace
 
COPY go.* .
RUN go mod download
COPY . .
RUN go mod tidy
