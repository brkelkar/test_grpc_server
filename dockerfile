FROM golang:1.15.6-alpine3.12

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/brkelkar/test_grpc_server

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .


# Install the package
RUN go build .
# This container exposes port 8080 to the outside world
EXPOSE 9000

# Run the executable
CMD ["./test_grpc_server"]