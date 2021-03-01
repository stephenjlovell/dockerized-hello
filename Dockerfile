FROM golang:1.16 as build
COPY dockerized-hello.go .
RUN go build dockerized-hello.go

FROM scratch
COPY --from=build /go/dockerized-hello /dockerized-hello
CMD ["/dockerized-hello"]