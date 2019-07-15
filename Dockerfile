FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /var/goapp/
RUN mkdir public
RUN mkdir templates
COPY roconfig .
COPY public ./public
COPY templates ./templates
COPY .env .
EXPOSE 7000
CMD ["./roconfig"]