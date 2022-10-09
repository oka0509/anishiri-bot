# Specify the base image for the go app.
FROM golang:latest
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/github.com/go-mysql
# Copy everything from this project into the filesystem of the container.
COPY . .
# Obtain the package needed to run redis commands. Alternatively use GO Modules.
RUN go mod tidy
# Compile the binary exe for our app.
RUN go build -o main .
# Start the application.
CMD ["./main"]

