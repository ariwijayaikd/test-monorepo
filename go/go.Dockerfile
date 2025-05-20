FROM golang:alpine AS builder

RUN apk update && \
    apk add --no-cache git && \
    apk add --no-cach bash && \
    apk add build-base \
    apk add curl \
    bash \
    make \
    ca-certificates \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY go.* ./
RUN go mod download
RUN go mod verify

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go .

########################## only take built application

FROM alpine:latest

RUN apk --no-cache add ca-certificates bash tzdata

ENV TZ=Asia/Jakarta

WORKDIR /app/

COPY --from=builder /app/go .

EXPOSE 8080

RUN ["chmod", "+x", "./go"]

CMD ["./go"]
