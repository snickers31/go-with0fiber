#  ./Dockerfile

FROM golang:1.19-bullseye AS builder

WORKDIR /build


COPY go.mod go.sum ./
RUN go mod download


COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM scratch

COPY --from=builder ["/build/apiserver", "/build/.env", "/"]

EXPOSE 5000

ENTRYPOINT ["/apiserver"]