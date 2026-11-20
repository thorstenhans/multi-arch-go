# multi-arch-go

A tiny Go HTTP API that gets built into a multi-arch Docker image (`linux/amd64` and `linux/arm64`) with `docker buildx`. It belongs to the blog post [How to Build Multi-Arch Docker Images with buildx](https://thorsten-hans.com/how-to-build-multi-arch-docker-images-with-ease/).

The API answers `GET /ping` with the CPU architecture of the image variant that serves the request, so you can see which platform actually runs.

The Dockerfile cross-compiles: the build stage always runs on the build machine's platform (`--platform=$BUILDPLATFORM`) and hands `TARGETOS`/`TARGETARCH` to the Go compiler, so no step runs under QEMU emulation.

## Build and run

Requires Docker with the containerd image store (the default in current Docker Desktop and Docker Engine) for `--load` of a multi-platform image.

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --tag multi-arch-go:dev \
  --load \
  .

docker run --rm -p 8080:8080 --platform linux/amd64 multi-arch-go:dev
curl localhost:8080/ping
# {"arch":"amd64","message":"pong"}
```

Swap `--platform linux/amd64` for `linux/arm64` to run the other variant.
