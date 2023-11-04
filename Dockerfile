FROM golang:1.21.0 AS builder
WORKDIR /app
COPY . .
RUN go build -o /app/build/main ./cmd

FROM nginx:latest
COPY --from=builder /app/build/main /usr/share/nginx/html
COPY --from=builder /app/static /usr/share/nginx/static
COPY --from=builder /app/templates /usr/share/nginx/templates
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["sh", "-c", "/usr/share/nginx/html/main & nginx -g 'daemon off;'"]
