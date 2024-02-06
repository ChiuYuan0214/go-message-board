FROM golang:latest as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN  go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest as release

WORKDIR /app

RUN apk --no-cache add ca-certificates libc6-compat

COPY --from=build /app/main .

EXPOSE 7080

CMD ["./main"]