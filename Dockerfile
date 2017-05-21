FROM alpine

ENV PORT 8080

EXPOSE $PORT

COPY currency-service /

CMD ["/bin/sh", "-c", "/bold-api"]
