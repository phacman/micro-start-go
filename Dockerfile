## Build
FROM golang:1.20-buster AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /micro-start-go

## Deploy
FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=build /micro-start-go ./app-start
USER nonroot:nonroot
ENTRYPOINT ["/app/app-start"]