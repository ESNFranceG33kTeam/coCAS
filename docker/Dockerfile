# Builder stage
FROM golang:1.19-alpine3.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp \
    HOME=/go \
    PROJECT_NAME=coCAS

LABEL org.opencontainers.image.title="${PROJECT_NAME}" \
      org.opencontainers.image.authors=${GITHUB_ACTOR} \
      org.opencontainers.image.vendor=${GITHUB_REPOSITORY} \
      org.opencontainers.image.source="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}" \
      org.opencontainers.image.url="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/tags" \
      org.opencontainers.image.description="Galaxy cas connector." \
      org.opencontainers.image.documentation="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/blob/main/README.md" \
      org.opencontainers.image.os="alpine" \
      org.opencontainers.image.version=${GITHUB_REF}

RUN apk add --no-cache make gcc curl jq

COPY . /src/
WORKDIR /src/

RUN make build-app

# Runner stage
FROM alpine:3.16 as runner

COPY --from=builder /src/bin/coCAS /usr/local/bin/coCAS

RUN adduser --uid 1000 --disabled-password --shell /bin/bash user
USER user

VOLUME /etc/coCAS/conf/
EXPOSE 8181

WORKDIR /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/coCAS" ]
CMD [ "--help" ]
