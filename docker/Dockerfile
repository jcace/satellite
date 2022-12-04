FROM golang:1.16

RUN mkdir /app
WORKDIR /app

COPY . .
RUN make build

CMD ["./artifacts/remote_desktop_ping_service_linux_amd64"]