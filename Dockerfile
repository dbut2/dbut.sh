FROM golang:alpine AS builder

COPY . /work

WORKDIR /work/go

RUN go build -o /bin/ssh ./cmd/ssh
RUN go build -o /bin/curl ./cmd/curl
RUN go build -o /bin/web ./cmd/web
RUN go build -o /bin/vanity ./cmd/vanity

FROM alpine:latest AS ssh

COPY --from=builder /bin/ssh /bin/ssh
COPY host_key host_key

EXPOSE 8022 8080
CMD ["/bin/ssh"]

FROM alpine:latest AS curl

COPY --from=builder /bin/curl /bin/curl

EXPOSE 8080
CMD ["/bin/curl"]

FROM alpine:latest AS web

COPY --from=builder /bin/web /bin/web

EXPOSE 8080
CMD ["/bin/web"]

FROM alpine:latest AS vanity

COPY --from=builder /bin/vanity /bin/vanity

EXPOSE 8080
CMD ["/bin/vanity"]