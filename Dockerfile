FROM golang:1.21.0

WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire application source code
COPY . .

# Build the application, specifying the path to the main.go file
RUN go build -v -o /usr/local/bin/app ./src/cmd/main.go

# Correct the CMD instruction to match the binary's location
CMD ["/usr/local/bin/app"]