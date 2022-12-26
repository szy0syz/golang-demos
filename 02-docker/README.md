# docker

- for dev

```dockerfile
FROM golang:1.19.2-bullseye

WORKDIR /app

COPY . .

# Install Go deps
RUN go mod download

# Builds your app with optional config
RUN go build -o/godokcer

EXPOSE 8080

# Specifies the executable command that runs when the container starts
CMD ["/godocker"]
```