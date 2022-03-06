FROM golang:alpine AS build
WORKDIR /httpserver/
#without this busybox can not work well
ENV CGO_ENABLED=0
COPY src/httpserver.go .
RUN go build -o httpserver httpserver.go

FROM busybox
WORKDIR /httpserver/
COPY --from=build /httpserver/httpserver /httpserver/httpserver
EXPOSE 80
CMD ["./httpserver"]