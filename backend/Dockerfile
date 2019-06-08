FROM golang:1.12.5-stretch AS builder
ENV GO111MODULE=on
COPY . /backend
WORKDIR /backend
RUN CGO_ENABLED=0 GOOS=linux go build -o backend

FROM alpine:latest
WORKDIR /root/
RUN apk add --no-cache tzdata git bash
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /backend .

RUN git clone https://github.com/vishnubob/wait-for-it.git
CMD ["./wait-for-it/wait-for-it.sh", "db:5432", "--", "./backend", "-usePostgres", "-resetPostgres", "-seedDatabase"]