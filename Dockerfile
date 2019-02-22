FROM golang AS builder
WORKDIR /go/src/github.com/eniot/release-proxy
COPY . .
RUN go get -d .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo

FROM scratch
WORKDIR /app
EXPOSE 80
COPY --from=builder /go/src/github.com/eniot/release-proxy/release-proxy .
ENTRYPOINT [ "/app/release-proxy" , "--addr", ":80"]