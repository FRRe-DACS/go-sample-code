FROM golang:1.12 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go get -v -d && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o hospital_server .
FROM scratch
COPY --from=builder /build/hospital_server /app/
COPY ./banner.txt /app/
WORKDIR /app
CMD ["./hospital_server"]