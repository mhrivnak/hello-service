# building container 
FROM registry.fedoraproject.org/fedora-minimal AS build
RUN mkdir -p /go/src && microdnf install golang && microdnf clean all 
RUN export GOPATH=/go
WORKDIR /go/src/
COPY hello.go /go/src/
RUN go build -o /bin/hello /go/src/hello.go

FROM registry.fedoraproject.org/fedora-minimal:30
COPY --from=build /bin/hello /usr/local/bin
CMD ["hello"]
