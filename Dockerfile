FROM golang:buster as builder

ENV GO111MODULE=on
LABEL maintainer="Eko Nur Rohman <ekonurrohman9x9@gmail.com>"

RUN apt update && apt install git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 
COPY ./ ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM golang:buster

RUN groupadd docker
RUN useradd -r -u 1001 -g docker eko9x9
USER eko9x9
WORKDIR /home/eko9x9/app

ENV GIN_MODE=release
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD [ "./main" ]