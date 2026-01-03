FROM golang

WORKDIR /portscanner

COPY go.mod .

RUN go mod tidy

COPY . .

# Build Executable file
RUN go build -o scanner main.go

# ENTRYPOINT
ENTRYPOINT ["./scanner"]