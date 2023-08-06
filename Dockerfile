# Build Stage: Build bot using the alpine image and goreleaser
FROM golang:1.20-alpine AS builder
RUN apk add --no-cache curl wget gnupg git upx
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o out/sangtei -ldflags="-w -s" .
RUN upx --brute out/sangtei

# Run Stage: Run bot using the bot binary copied from build stage
FROM alpine
COPY --from=builder /app/out/sangtei /app/sangtei
ENTRYPOINT ["/app/sangtei"]

LABEL org.opencontainers.image.authors="Nicky Lalrochhara <nickylalrochhara917@gmail.com>"
LABEL org.opencontainers.image.url="https://github.com/lalrochhara/Sangtei"
LABEL org.opencontainers.image.source="https://github.com/lalrochhara/Sangtei"
LABEL org.opencontainers.image.title="Miss Sangtei Go Docker Image"
LABEL org.opencontainers.image.description="Miss Sangtei Go Docker Image"
LABEL org.opencontainers.image.vendor="NickyLrca"
