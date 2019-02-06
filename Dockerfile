FROM golang:alpine as builder
WORKDIR /build
COPY main.go ./
RUN go build -o ./main .

FROM alpine 
EXPOSE 8080
WORKDIR app
COPY --from=builder /build/main ./
CMD ["./main"]
