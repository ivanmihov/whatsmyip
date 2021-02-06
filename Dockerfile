FROM golang:alpine as build
ARG ARCH=amd64
COPY . /opt
WORKDIR /opt
RUN CGO_ENABLED=0 GOARCH=$ARCH GOOS=linux go build  -o ./whatsmyip

FROM scratch
WORKDIR /opt
COPY --from=build /opt/whatsmyip .
EXPOSE 80
ENTRYPOINT ["/opt/whatsmyip"]
