FROM golang:latest
LABEL author="Edward Pie"
ENV FFPUPDATER_PORT=9000
ENV FFPUPDATER_ENV=development
ENV FFPUPDATER_SYNC_FREQUENCY=1
ENV FFPUPDATER_DATABASE_URI=""
ENV SRC_DIR=/go/src/github.com/hackstock/ffp-updater-service
ADD . ${SRC_DIR}
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR ${SRC_DIR}
RUN dep ensure -v
RUN go build -race .
ENTRYPOINT [ "./ffp-updater-service" ]