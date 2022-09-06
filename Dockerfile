FROM golang:alpine

Run mkdir -p /var/www

WORKDIR /var/www

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /var/www

RUN go build -o app

EXPOSE 8080

ENTRYPOINT ["/var/www/app"]
