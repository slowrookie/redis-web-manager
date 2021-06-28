FROM alpine

ENTRYPOINT ["/usr/bin/RWM"]

COPY RWM /usr/bin/RWM

EXPOSE 9090