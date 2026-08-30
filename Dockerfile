FROM golang:1.26-trixie

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o /app/exe cmd/the_wall/main.go

CMD [ "/app/exe" ]
