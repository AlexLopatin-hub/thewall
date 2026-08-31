FROM golang:1.26-trixie

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go build -o /app/exe ./cmd/thewall/main.go

EXPOSE 8080
CMD [ "/app/exe" ]
