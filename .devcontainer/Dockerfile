ARG GO_VERSION=1.18
ARG ALPINE_VERSION=3.10

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

# Install packages and Go language server
RUN apk add -q --update --progress --no-cache git sudo openssh-client zsh
RUN go get -u -v golang.org/x/tools/cmd/gopls 2>&1

RUN apt update \
    && apt install -y bash \
                   build-essential \
                   git \
                   curl \
                   ca-certificates \
                   wget \
                   zip \
                   awscli \
    && rm -rf /var/lib/apt/lists