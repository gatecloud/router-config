FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /var/goapp
RUN mkdir public
RUN mkdir templates
COPY public ./public
COPY templates ./templates
COPY nginx.conf /etc/nginx/conf.default.conf
COPY roconfig .
COPY .env .
EXPOSE 7000 
CMD ["./roconfig"]