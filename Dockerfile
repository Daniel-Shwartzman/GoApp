FROM golang:1.19

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. 
COPY *.go ./

# Copy the HTML file
COPY index.html ./

# Run the tests
RUN go test -v ./...

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app

EXPOSE 8081

# Run
CMD ["/app/app"]
