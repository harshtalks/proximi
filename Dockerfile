# DockerFile for the Proximi App

FROM golang:alpine

# Making work directory

WORKDIR /App

# Copying everything

COPY . .

# Getting go files

RUN go get

# ENV Arguments

ARG JWT_SECRET
ARG DATABASE_HOST
ARG DATABASE_PASSWORD
ARG DATABASE_PORT
ARG DATABASE_NAME
ARG DATABASE_USER
ARG BING_KEY
ARG BING_URL

# BUILD

RUN go build -o bin .

# RUN

ENTRYPOINT [ "/App/bin" ]
