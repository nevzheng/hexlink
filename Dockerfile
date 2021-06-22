ARG port=8080 # Which port to expose

##############
# BUILD STAGE
##############
FROM golang:alpine AS builder

# Set necessary environment variables for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    COOS=linux \
    GOARCH=amd64

# Move to working directory
WORKDIR /build

# Pull dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Pull source
COPY . .

# Build hexlink
RUN go build -o hexlink-server ./cmd/hexlink-server/main.go

# Move to /dist for resulting binary
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/hexlink-server .

##############
# DEPLOY STAGE
##############
FROM scratch

COPY --from=builder /dist/hexlink-server /

# Command to use when starting the container
ENTRYPOINT ["/hexlink-server"]
