FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o weather ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/weather .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD [ "./weather" ]
