FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY ./processor ./processor
RUN go build -o app processor/cmd/main.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 7259
CMD [ "./app" ]
