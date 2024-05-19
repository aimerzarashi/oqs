FROM golang:1.22.1-alpine3.19 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd ./
COPY internal ./

WORKDIR /app/main
RUN go build -o /docker-exec

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-exec /docker-exec

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT ["/docker-exec"]