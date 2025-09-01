FROM golang:1.24.1 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o quiz-api cmd/api/main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=build /app/quiz-api ./
ENTRYPOINT ["./quiz-api"]