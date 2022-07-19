# Go + Alpine so we can compile
FROM golang:1.18.4-alpine3.16 AS build

WORKDIR /tmp/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o eviecoin bin/main.go

# Switch to alpine for a lightweight image since we don't need to install anything, native executable go brr
FROM alpine

COPY --from=build /tmp/app/eviecoin /bin/eviecoin

ENTRYPOINT ["/bin/eviecoin"]