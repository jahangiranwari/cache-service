FROM golang:1.16-buster

WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o ./ ./...

ARG APP_ENV
RUN echo "$APP_ENV"
RUN if [ "$APP_ENV" = "development" ]; then \
      go get github.com/go-delve/delve/cmd/dlv \
      github.com/uudashr/gopkgs/v2/cmd/gopkgs \
      github.com/ramya-rao-a/go-outline \
      github.com/cweill/gotests/gotests \
      github.com/fatih/gomodifytags \
      github.com/josharian/impl \
      github.com/haya14busa/goplay/cmd/goplay \
      honnef.co/go/tools/cmd/staticcheck  \
      golang.org/x/tools/gopls; \
    fi

# Install redis-cli (for dev)
RUN if [ "$APP_ENV" = "development" ]; then \
      cd /tmp \
      && wget http://download.redis.io/redis-stable.tar.gz \
      && tar xvzf redis-stable.tar.gz \
      && cd redis-stable \
      && make \
      && cp src/redis-cli /usr/local/bin/ \
      && chmod 755 /usr/local/bin/redis-cli \
      && rm -rf /tmp/redis-stable.tar.gz \
      && rm -rf /tmp/redis-stable; \
    fi
