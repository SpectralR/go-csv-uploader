FROM golang:1.26-alpine

WORKDIR /app

ENV PATH="/go/bin:${PATH}"

RUN go install github.com/air-verse/air@latest

COPY go.mod ./

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]