FROM golang:1.18.4-bullseye AS builder

LABEL maintainer="neusta schweiz"

# Install dependencies to be able to build our code
#RUN apk add --no-cache --update git && apk add build-base

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY ./src/*.go ./

# Download all required packages
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./src .

# Compile the code

RUN go build -o output/server main.go


FROM golang:1.18.4-bullseye AS final
WORKDIR /app
COPY --from=builder app/output/server .
EXPOSE 80
EXPOSE 443
#RUN groupadd -r neusta && useradd --no-log-init -r -g neusta neusta
#RUN chown -R neusta:neusta  /app
#USER neusta
ENV GIN_MODE=release
CMD ["./server"]