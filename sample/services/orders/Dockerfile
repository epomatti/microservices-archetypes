FROM golang AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build

FROM scratch AS runner
COPY --from=builder /app/orders /app/orders
ENTRYPOINT ["/app/orders"]