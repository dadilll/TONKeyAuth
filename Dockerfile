FROM golang:1.24.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -o OauthTON ./cmd/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/OauthTON .

COPY conf /app/conf
COPY key /app/key

CMD ["./OauthTON"]