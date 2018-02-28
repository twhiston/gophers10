FROM tomwhiston/micro-golang:toolchain

ARG ARCH="x86_64"
ARG DOCKER_CHANNEL=edge
ARG DOCKER_VERSION=17.11.0-ce

RUN apk add --no-cache -U bash curl && \
    mkdir -p /root/.docker/ && \
    curl -fL -o docker.tgz "https://download.docker.com/linux/static/${DOCKER_CHANNEL}/${ARCH}/docker-${DOCKER_VERSION}.tgz" && \
    tar --extract \
        --file docker.tgz \
        --strip-components 1 \
        --directory /usr/local/bin/ && \
    rm docker.tgz && \
    gometalinter --install

ADD https://raw.githubusercontent.com/aelsabbahy/goss/v0.3.5/extras/dgoss/dgoss /usr/local/bin/dgoss
RUN chmod +rx /usr/local/bin/dgoss