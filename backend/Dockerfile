#
# build stage
#
FROM golang:1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /app/knoq

#
# runtime stage
#
FROM gcr.io/distroless/static-debian11:latest

WORKDIR /app

COPY --from=builder /app/knoq /app/

CMD ["/app/knoq"]