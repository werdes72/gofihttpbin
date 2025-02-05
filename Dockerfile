from golang:1.24rc2 AS build

WORKDIR /go/src/gofihttpbin
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gofihttpbin ./cmd/gofihttpbin/gofihttpbin.go


from alpine:3.21.2

WORKDIR /gofihttpbin
RUN mkdir -p ./web/static
COPY ./web/static ./web/static

COPY --from=build /go/src/gofihttpbin/gofihttpbin .

RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["./gofihttpbin"]
