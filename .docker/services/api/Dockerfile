FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./bin/api ./cmd/api/main.go

#FROM scratch
#COPY --from=0 /bin/hello /bin/hello
EXPOSE 8080
CMD ["./bin/api"]


