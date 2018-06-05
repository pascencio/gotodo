FROM golang:1.10-alpine

ARG http_proxy

COPY run.sh /usr/bin/gotodo

RUN export http_proxy=${http_proxy} && \
export https_proxy=${http_proxy} && \
export HTTP_PROXY=${http_proxy} && \
export HTTPS_PROXY=${http_proxy} && \
apk add --update inotify-tools curl && \
chmod +x /usr/bin/gotodo && \
mkdir -p /go/src/github.com/pascencio/gotodo && \
chmod 777 -R /go/src/github.com && \
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
export http_proxy="" && \
export https_proxy="" && \
export HTTP_PROXY="" && \
export HTTPS_PROXY=""

WORKDIR /go/src/github.com/pascencio/gotodo

EXPOSE 8080

CMD [ "/usr/bin/gotodo" ]