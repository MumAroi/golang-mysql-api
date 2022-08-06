# base image
FROM golang:alpine as builder

# maintainer 
LABEL maintainer="paramas.wae.th@gmail.com"

# install git
RUN apk update && apk add --no-cache git

# working directory 
WORKDIR /app

# copy go mod and sum
COPY go.mod go.sum ./

# download go dependencies
RUN go mod download 

#  copy source code
COPY . .

# build go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# start a new stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# working directory 
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .     

# Expose port 8080 to the outside world
EXPOSE 8080

# executable
CMD ["./main"]
