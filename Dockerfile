FROM golang:1.22-alpine AS build
WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o /bin/email-verifier ./cmd/main.go

FROM alpine:latest
COPY --from=build /bin/email-verifier /usr/local/bin/email-verifier

EXPOSE 8080

ENTRYPOINT [ "email-verifier" ]
