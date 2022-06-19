FROM --platform=linux/x86_64 golang:1.18.3-alpine as builder

ENV CGO_ENABLED 0

RUN go env -w GOPROXY=direct

WORKDIR /biuld

RUN apk add --update && apk add git

COPY go.mod go.sum ./
RUN go mod download

ENV GOBIN /build/go/bin
COPY go go
RUN go install ./go/...

FROM --platform=linux/x86_64 golang:1.18.3-alpine

COPY --from=builder /build/go/bin/hello /usr/local/bin/hello
ENTRYPOINT [ "/usr/local/bin/preprocess" ]
