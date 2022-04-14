FROM golang:1.18-alpine

WORKDIR /swapknowhow/courses

COPY go.mod ./

RUN go mod tidy

EXPOSE 8084

COPY . .
COPY /api ./api
COPY /internal ./internal
COPY *.go ./

RUN go build

CMD ["./courses"]