# ARG baseImage="golang:1.18-alpine"
FROM golang:1.18-alpine as builder
# ARG HEADER_FILE
# ARG ENV_FILE
# ENV HEADER_FILE=$HEADER_FILE
# ENV ENV_FILE=$ENV_FILE

# RUN echo "File swagger: $HEADER_FILE"
# RUN echo "File env: $ENV_FILE"

RUN apk add bash ca-certificates git gcc g++ libc-dev

# Here we copy the rest of the source code
RUN mkdir -p /projects/phenikaa-embedded
WORKDIR /projects/phenikaa-embedded

# We want to populate the module cache based on the go.{mod,sum} files. 
COPY go.mod .
COPY go.sum .
RUN ls -la /projects/phenikaa-embedded

RUN go mod download

# COPY $HEADER_FILE /projects/pdt-phenikaa-htdn-backend/$HEADER_FILE

COPY . /projects/phenikaa-embedded
# COPY .env.pro /projects/phenikaa-embedded/.env

# RUN go install github.com/swaggo/swag/cmd/swag@v1.8.4
# RUN swag init --parseDependency -g $HEADER_FILE

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o phenikaa-embedded .