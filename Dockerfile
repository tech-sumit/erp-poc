FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git
WORKDIR /app

COPY . /app/

RUN echo "[url \"https://sumit-agrawal:Xqnft4nZcfs5KujE83m6@bitbucket.org/perennialsys\"] insteadOf = https://bitbucket.org/perennialsys" >> /root/.gitconfig

ENV GOPRIVATE="bitbucket.org/perennialsys"
RUN go mod download

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o main.go -o build/service

FROM alpine as release
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/build/service /service
EXPOSE ${SERVICE_HOST}

ENTRYPOINT ["./service"]
