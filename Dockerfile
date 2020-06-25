

ARG CMD_NAME=tfto


# # # ============================================================


FROM golang:1.14.4-alpine3.12 AS compile

# use global ARG (top of Dockerfile)
ARG CMD_NAME

RUN apk add git --no-cache

ENV GO111MODULE on
ENV WORKDIR /workdir/
WORKDIR ${WORKDIR}

COPY go.mod go.sum ${WORKDIR}
RUN go mod download

COPY . ${WORKDIR}
RUN go install . || go install ./cmd/${CMD_NAME}/


# # # ============================================================


FROM alpine:3.12

# use global ARG (top of Dockerfile)
ARG CMD_NAME

ENV CMD_NAME ${CMD_NAME}

COPY --from=compile /go/bin/${CMD_NAME} /${CMD_NAME}
RUN chmod +x /${CMD_NAME}

# ENTRYPOINT ["/hello"]

# notice: exec form does not invoke a command shell => we need the `sh` wrapper
# ENTRYPOINT ["sh", "-c", "/${CMD_NAME}"]

# use `shell form`
ENTRYPOINT /${CMD_NAME}

