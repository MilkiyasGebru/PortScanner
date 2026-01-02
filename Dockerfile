FROM golang

WORKDIR /portscanner

COPY . .

RUN go mod tidy

# Build Executable file
RUN go build -o scanner main.go

# ENTRYPOINT
ENTRYPOINT ["./scanner"]