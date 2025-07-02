FROM golang:1.21.6-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY ./user/. ./              # User service code
COPY ./proto ./proto          # Shared proto folder

RUN go mod download
RUN go build -o /app/user-service ./cmd/main.go

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache postgresql-client

COPY --from=builder /app/user-service .
COPY user/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 50051
CMD ["/wait-for-it.sh", "postgres", "5432", "--", "./user-service"]
