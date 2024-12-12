FROM golang:1.23-bookworm

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

EXPOSE 3030

CMD ["/backend/api"]
