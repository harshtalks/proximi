# DockerFile for the Proximi App

FROM golang:alpine as builder

# Making work directory

WORKDIR /App

# Copying everything

COPY . .

# COPY ENV

# Getting go files

RUN go get

# BUILD

RUN go build -o main .

# Stage Building

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /App/main .

EXPOSE 8080

# RUN

CMD [ "./main" ]
