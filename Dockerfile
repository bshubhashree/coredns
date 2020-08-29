FROM golang:1.14
WORKDIR /coredns
ADD . /coredns
RUN cd /coredns && go build -o coredns
EXPOSE 53 53/udp
ENTRYPOINT ./coredns

# FROM scratch
# COPY --from=0 /coredns /coredns

# EXPOSE 53 53/udp
# CMD ["/coredns"]
