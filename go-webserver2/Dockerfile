FROM golang:1.16-alpine AS builder

RUN mkdir /build/
WORKDIR /build/
COPY go.* ./
RUN go mod download
COPY . ./
WORKDIR /build/
RUN go build -o webapp

FROM alpine
WORKDIR /app
EXPOSE 8080
COPY --from=builder /build/webapp /
ENTRYPOINT /webapp