FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM scratch
ENV MONGODB_USERNAME=testUser MONGODB_PASSWORD=testPassword
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
EXPOSE 8080