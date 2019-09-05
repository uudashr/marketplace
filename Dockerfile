# == Builder ==
FROM golang:1.13.0-alpine3.10 as builder

RUN apk add --no-cache bash=5.0.0-r0 git=2.22.0-r0

WORKDIR /app

# Copy dependencies definition
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy remaining source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /app/marketplace-up cmd/marketplace-up/*.go

# == Runner ==
FROM alpine:3.10.1

# Copy binary from builder
COPY --from=builder /app/marketplace-up /app/marketplace-up

CMD ["/app/marketplace-up"]