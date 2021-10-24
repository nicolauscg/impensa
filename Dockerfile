FROM golang:1.14-alpine AS build

WORKDIR $GOPATH/src/github.com/nicolauscg/impensa
RUN apk update --quiet && apk add --quiet ca-certificates
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /impensa-be .

FROM alpine:3.11
WORKDIR /app/
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /impensa-be /app/
COPY .env /app/
COPY conf/app.conf /app/conf/
CMD ["./impensa-be"]
