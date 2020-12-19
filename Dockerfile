FROM busybox:latest

WORKDIR /app

COPY ./registry-auth .

ENTRYPOINT ["/app/registry-auth"]
