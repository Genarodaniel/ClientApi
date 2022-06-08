FROM golang:1.17.3-alpine3.14
WORKDIR /api

ENV GO111MODULE=on CGO_ENABLED=0
EXPOSE 80
ENV PORT 80
ENV HTTP_PORT 80
ENV HTTP_ADDRESS 0.0.0.0
ENV TZ America/Sao_Paulo

COPY go.mod go.sum /api/
RUN go mod download

COPY . .
RUN go build -o /api/main /api/main.go

CMD [ "/api/main" ]