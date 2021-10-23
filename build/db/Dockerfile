FROM golang:bullseye AS build

# Enable the vendoring mode
ENV GOFLAGS="-mod=vendor"

ADD . /dockerdev
WORKDIR /dockerdev

RUN CGO_ENABLED=0 cd cmd/db && go build -o /db

EXPOSE 8083
EXPOSE 8084

# Final stage
FROM debian:latest

WORKDIR /
COPY --from=build /db /

CMD ["/db"]