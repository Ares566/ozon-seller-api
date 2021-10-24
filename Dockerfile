#Полный образ golang содержит все чтоб билдить,
#но он очень большой, поэтому билдим тут...
FROM golang:latest AS buildContainer

WORKDIR /go/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 go build -o ./ozoneapi ./cmd/main.go

# ...а потом выносим в финальный легкий контейнер
FROM alpine:latest


RUN mkdir -p /app

WORKDIR /app
COPY --from=buildContainer /go/src/app/ozoneapi .
COPY --from=buildContainer /go/src/app/.env .

# обязательно раскомментить перед продом
#ENV GIN_MODE release

# настройки лежат в .env
#ENV CACHE_DRIVER REDIS
#ENV CACHE_HOST redis
#ENV CACHE_PORT 6379


CMD ["./ozoneapi"]

