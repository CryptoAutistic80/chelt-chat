FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY backend/go.mod ./
# Initialize go.sum if it doesn't exist
RUN go mod download
COPY backend/ .
RUN go build -o chelt-chat-backend

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/chelt-chat-backend .
EXPOSE 8080
CMD ["./chelt-chat-backend"] 