FROM golang:bullseye AS build

# Enable the vendoring mode
ENV GOFLAGS="-mod=vendor"

ADD . /dockerdev
WORKDIR /dockerdev

RUN CGO_ENABLED=0 cd cmd/auth && go build -o /auth

EXPOSE 8081
EXPOSE 8082
EXPOSE 8084

# Final stage
FROM debian:latest

WORKDIR /
COPY --from=build /auth /

CMD ["/auth"]