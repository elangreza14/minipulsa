# Build stage
FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod . 
# <- COPY go.mod and go.sum files to the workspace
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:latest
RUN apk add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/app.env ./app.env

EXPOSE 9001
CMD [ "/app/main" ]