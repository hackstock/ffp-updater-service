FROM golang:latest
LABEL author="Edward Pie"
ENV SRC_DIR=/ffp-updater-service
ADD . ${SRC_DIR}
WORKDIR ${SRC_DIR}
RUN go build -race .
ENTRYPOINT [ "./ffp-updater-service" ]