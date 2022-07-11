FROM golang:1.16-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go build -o /go/bin/ayr cmd/ayr/main.go

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/ayr /
CMD ["./ayr"]