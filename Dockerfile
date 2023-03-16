FROM golang:alpine

RUN apk update && apk --nocache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]