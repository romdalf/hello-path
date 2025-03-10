# hello-path


## deploy with community

Adapt to match your needs!

```ini
FROM golang:1.20 

WORKDIR /app

COPY go.mod ./
RUN go mod edit -module=github.com/romdalf/hello-path
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /hello

EXPOSE 8080

ENTRYPOINT ["/hello"]
```

```
podman build -t ghcr.io/romdalf/hello:1.1-pub -f Containerfile.pub
```

## deploy with ubi

Adapt to match your needs!

```ini
FROM registry.access.redhat.com/ubi8/go-toolset AS build

# prefer the SHA tagging like registry.access.redhat.com/ubi8/go-toolset@sha256:168ac23af41e6c5a6fc75490ea2ff9ffde59702c6ee15d8c005b3e3a3634fcc2 AS build

RUN go mod init hello
RUN go mod tidy
#RUN go mod download 

COPY *.go ./

RUN CGO_ENABLED=0 go build -o .

FROM registry.access.redhat.com/ubi8/ubi-micro

# prefer the SHA tagging registry.access.redhat.com/ubi8/ubi-micro@sha256:6a56010de933f172b195a1a575855d37b70a4968be8edb35157f6ca193969ad2

COPY --from=build ./opt/app-root/src/hello .

EXPOSE 8080
ENTRYPOINT ["./hello"]
```

```
podman build -t ghcr.io/romdalf/hello:1.1-ubi -f Containerfile.ubi
```
