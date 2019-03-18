FROM golang:1.8-alpine
COPY catweb.go /go
RUN go build catweb.go

FROM alpine:latest
COPY --from=0 /go/catweb /
COPY templates/index.html /
COPY static/* /static/
EXPOSE 8080
CMD /catweb
