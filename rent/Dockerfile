#FROM golang:1.22.4-alpine3.20 as modules
#COPY go.mod go.sum /modules/
#WORKDIR /modules
#RUN go mod download
#
#FROM golang:1.22.4-alpine3.20 as builder
#
#COPY --from=modules /go/pkg /go/pkg
#COPY . /app
#COPY .env /app
#WORKDIR /app
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
#    go build -tags migrate -o /bin/app ./cmd/app
#
#FROM scratch
#COPY --from=builder /app/config /config
#COPY --from=builder /app/migrations /migrations
#COPY --from=builder /bin/app /app
#
##WORKDIR /app
##
##COPY go.mod ./
##RUN go mod download
##
###COPY *.go ./
##COPY *.* ./
##
##RUN #go build -o ./cmd/app
#
#EXPOSE 8080
#
#CMD [ "/app" ]


# Use the official Golang image as the base image
FROM golang:1.22.4-alpine3.20

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main ./cmd/app/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]