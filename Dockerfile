FROM golang:1.17-stretch

COPY cmd/api/config.yaml golang-projects-a/cmd/api/config.yaml
COPY main golang-projects-a/main

EXPOSE 8080
WORKDIR "golang-projects-a"
ENTRYPOINT ["./main"]