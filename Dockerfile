
FROM harbor.myshuju.top/gcr.io/distroless/static:nonroot
WORKDIR /
COPY /workspace/app .
EXPOSE 8080

ENTRYPOINT ["./app"]