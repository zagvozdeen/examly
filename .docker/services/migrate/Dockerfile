FROM golang as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/migrate ./cmd/migrate/main.go

FROM alpine

WORKDIR /app

COPY --from=build /app/bin/migrate .
COPY --from=build /app/.env .
COPY --from=build /app/migrations ./migrations/

ENTRYPOINT ["/bin/sh"]


