FROM golang as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o /app/bin/api ./cmd/api/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/bin/api .
COPY --from=build /app/.env .

EXPOSE 8000

CMD ["./api"]


