FROM alpine

ENTRYPOINT ["/usr/bin/RWM"]

COPY RWM /usr/bin/RWM

VOLUME [ "/rwm_data" ]

EXPOSE 9090