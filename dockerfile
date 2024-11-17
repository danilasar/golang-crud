FROM golang:latest AS build
WORKDIR /app
COPY ./src/go.mod ./
RUN go mod download
# RUN go mod tidy -v
COPY ./src/ ./
RUN go build -o output/server .

## Deploy
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/configs /configs
COPY --from=build /app/output/server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
