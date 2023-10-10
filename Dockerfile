FROM golang:1.20 as BUILD

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -o /main

FROM gcr.io/distroless/base

WORKDIR /

COPY --from=build /main /main

COPY . /

EXPOSE 5000

USER root

ENTRYPOINT ["/main"]