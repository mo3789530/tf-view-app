FROM golang:1.19-alpine as dev

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update
RUN apk add make git
RUN apk add nodejs npm
COPY go.mod go.sum ./
COPY . ${ROOT}

RUN export NODE_OPTIONS=--openssl-legacy-provider
RUN make build-ui
RUN go mod download

CMD ["go", "run", "main.go"]

FROM golang:1.19-alpine as builder

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update 
RUN apk add make git
RUN apk add npm nodejs

COPY go.mod go.sum ./
COPY . ${ROOT}

RUN export NODE_OPTIONS=--openssl-legacy-provider
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux make build


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/app/binary"]
