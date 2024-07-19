FROM golang:1.22-alpine AS build_base

#RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/racechip

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
COPY config.yaml .

RUN go mod download

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/racechip .

# Start fresh from a smaller image
FROM alpine:3.14
#RUN apk add ca-certificates

COPY --from=build_base /tmp/racechip/out/racechip /app/racechip
COPY --from=build_base /tmp/racechip/config.yaml /app/config.yaml
# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
#CMD ["/app/racechip"]
WORKDIR /app
ENTRYPOINT ["/app/racechip"]