
FROM golang:1.18-bullseye
#RUN go install github.com/spf13/viper@latest


ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV CONFIG_VALUE="hello there from docker env variable"
ENV APP_HOME /go/src/configurablevariable

RUN mkdir -p "${APP_HOME}"
ADD src "/go/src/configurablevariable"

#RUN go get github.com/spf13/viper


WORKDIR "/go/src/configurablevariable"
RUN go mod vendor
RUN go install
RUN go build -o main
EXPOSE 10000
CMD ["/go/src/configurablevariable/main"]