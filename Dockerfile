FROM golang:1.12.6-alpine3.10 AS compile
RUN apk add git --no-cache

ENV GO111MODULE on
ENV WORKDIR /go/src/github.com/pandorasnox/tftoinv/
WORKDIR ${WORKDIR}

COPY go.mod go.sum ${WORKDIR}
RUN go mod download

COPY . ${WORKDIR}
RUN go install .

# # # ============================================================
FROM alpine:3.8
COPY --from=compile /go/bin/tftoinv /tftoinv
ENTRYPOINT ["/tftoinv"]
