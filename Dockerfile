# Running Xpense in Docker is Experimental - not recommended for production use!

# Example of usage:
# docker build -t xpense .
# docker run --name xpense1 --entrypoint xpensetool xpense --datadir=/var/xpense genesis fake 1
# docker run --volumes-from xpense1 -p 5050:5050 -p 5050:5050/udp -p 18545:18545 xpense --fakenet 1/1 --http --http.addr=0.0.0.0

FROM golang:1.22 as builder

RUN apt-get update && apt-get install -y git musl-dev make

WORKDIR /go/Xpense
COPY . .

ARG GOPROXY
RUN go mod download
RUN make all


FROM golang:1.22

COPY --from=builder /go/Xpense/build/xpensed /usr/local/bin/
COPY --from=builder /go/Xpense/build/xpensetool /usr/local/bin/

EXPOSE 18545 18546 5050 5050/udp

VOLUME /var/xpense

ENTRYPOINT ["xpensed", "--datadir=/var/xpense"]
