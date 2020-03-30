FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/github.com/nicolauscg/impensa
COPY . ./
COPY .env.* /impensa-be/
RUN go build -o /impensa-be/impensa

FROM alpine:3.11
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /impensa-be/ /app/
CMD ["./impensa"]
