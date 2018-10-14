FROM golang:latest
LABEL author="Edward Pie"
ENV SRC_DIR=/go/src/github.com/hackstock/ffp-updater-service
ADD . ${SRC_DIR}
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR ${SRC_DIR}
RUN dep ensure -v
RUN go build -race .
ENTRYPOINT [ "./ffp-updater-service" ]