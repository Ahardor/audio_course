FROM golang:alpine AS builder
WORKDIR /app
COPY . ./
RUN go build -o app mock/cmd/main.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 7249
CMD [ "./app" ]
