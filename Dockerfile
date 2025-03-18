FROM golang:1.24

# set working directory
WORKDIR /go/src/app

# copy source files
COPY . .

# EXpose the port
EXPOSE 8080

# build the app
RUN go build -o main cmd/main.go

# run the app
CMD ["./main"]