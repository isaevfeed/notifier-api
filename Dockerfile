FROM golang:alpine

WORKDIR /go/bin

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o ./notifier ./cmd/server/main.go

ENTRYPOINT ["./notifier"]
