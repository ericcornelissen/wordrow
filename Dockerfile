FROM golang:1.15-alpine

# Install development tools
RUN apk add --no-cache \
  gcc \
  git \
  libc-dev \
  make

# Install NodeJS for npx
RUN apk add --no-cache \
  nodejs \
  npm

# Set up
WORKDIR /go/src/wordrow
COPY Makefile ./
RUN make install

# Remove build-only tools
RUN apk del \
  gcc \
  libc-dev

# Add project
COPY . .
