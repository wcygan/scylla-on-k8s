# First stage: Build the Protobuf files using Buf
FROM bufbuild/buf:latest as bufbuild
WORKDIR /app
COPY buf.gen.yaml ./
COPY generated ./generated
COPY proto ./proto
RUN buf generate proto

# Second stage: Download dependencies
FROM golang:1.22 as godeps
WORKDIR /app
COPY --from=bufbuild /app/generated /app/generated
WORKDIR /app/server
COPY server/go.mod server/go.sum ./
RUN go mod download

# Third stage: Build the Go binary
FROM golang:1.22 as gobuild
WORKDIR /app

# Copy the generated files from the previous stage
COPY --from=bufbuild /app/generated /app/generated

# Copy the downloaded dependencies from the second stage
COPY --from=godeps /go/pkg /go/pkg

# Adjust the WORKDIR to the specific directory where the main.go is located if it's not in the root
WORKDIR /app/server

# Copy the rest of the code
COPY server/ ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ../build/server cmd/main.go

# Final stage: Create the runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the compiled binary from the third stage
COPY --from=gobuild /app/build/server ./server

CMD ["./server"]