FROM golang:1.13 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.9
WORKDIR /app
COPY --from=build /app/main .
CMD ["./main"]
