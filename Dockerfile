FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /bin/ecr-docker-password ./cmd/ecr-docker-password

FROM alpine:3.18
COPY --from=builder /bin/ecr-docker-password /bin/ecr-docker-password
ENTRYPOINT ["/bin/ecr-docker-password"]