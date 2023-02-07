FROM golang:1.18.1 as golang-build

WORKDIR /app-build

COPY . .

RUN ls -alt && go build -o hello app.go

RUN pwd

FROM alpine

RUN apk add --no-cache libc6-compat

COPY --from=golang-build /app-build/hello /usr/local/bin/hello
COPY --from=golang-build /app-build/static ./static
RUN chmod +x /usr/local/bin/hello

CMD ["/usr/local/bin/hello"]

ENTRYPOINT [ "/bin/sh", "-l", "-c" ]