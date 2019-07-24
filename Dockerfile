FROM nginx:alpine
RUN apk --no-cache add ca-certificates
WORKDIR /var/goapp
RUN mkdir public
RUN mkdir templates
COPY roconfig .
COPY public ./public
COPY templates ./templates
COPY nginx.conf /etc/nginx/conf.default.conf
COPY .env .
EXPOSE 80
# CMD ["sh", "-c", "service nginx start && ./roconfig"]