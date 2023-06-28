FROM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache gcc musl-dev
CMD mkdir /app
WORKDIR /app
COPY . .
RUN go get github.com/gofiber/fiber/v2
RUN go get gorm.io/gorm 
RUN go get github.com/mattn/go-sqlite3

RUN go build -o main .

FROM alpine:latest
CMD mkdir /app
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/view ./view

EXPOSE 8080
ENTRYPOINT ["./main"]



