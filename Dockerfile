FROM golang:1.23.5-alpine as build

WORKDIR /work/MiniHttpReMux
COPY . .

RUN go mod tidy
RUN go mod download
RUN go vet -v 
RUN go test -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app 

#FROM gcr.io/distroless/static-debian12
FROM scratch

VOLUME ["/config"]

COPY --from=build /work/MiniHttpReMux/bin/app /
COPY --from=build /work/MiniHttpReMux/config /config
CMD ["/app"]