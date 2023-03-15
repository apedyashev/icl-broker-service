FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD ["/app/brokerApp"]

# commented code below is needed if makefile is not used
########################################################################


# #base go image
# FROM golang:1.18-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# # build a tiny docker image
# FROM alpine:latest

# RUN mkdir /app

# COPY --from=builder /app/brokerApp /app

# CMD ["/app/brokerApp"]