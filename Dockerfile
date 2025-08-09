FROM golang:tip-alpine3.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM gcr.io/distroless/static-debian12
COPY --from=builder /main /
CMD ["/main"]
