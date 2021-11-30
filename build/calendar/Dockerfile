FROM golang:bullseye AS build

# Enable the vendoring mode
ENV GOFLAGS="-mod=vendor"

ADD . /dockerdev
WORKDIR /dockerdev

RUN CGO_ENABLED=0 cd cmd/calendar && go build -o /calendar

EXPOSE 8081
EXPOSE 8082
EXPOSE 8084

# Final stage
FROM debian:latest

COPY --from=golang:bullseye /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

WORKDIR /
COPY --from=build /dockerdev /calendar-app
COPY --from=build /calendar /

CMD ["/calendar"]