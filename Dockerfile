FROM golang as builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 go build -o app .

FROM alpine
ENV BARK_KEY=""
ENV INTERVAL=60
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/holdon.txt .
CMD ["/app/app"]