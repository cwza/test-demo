FROM alpine:latest
USER root

COPY ./main /main
RUN chmod +x /main

CMD ["/main"]