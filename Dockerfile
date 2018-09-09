FROM golang:latest
LABEL author="Edward Pie"
ENV SMS_API_USERNAME="AWA"
ENV SMS_API_PASSWORD="ttvpass101"
ENV SMS_API_SENDERID="FlyAWA"
ENV FFPUPDATER_PORT=9000
ENV FFPUPDATER_ENV=development
ENV FFPUPDATER_SYNC_FREQUENCY=1
# ENV FFPUPDATER_DATABASE_URI=mysql://flyawa:Nf@2el3w@35.176.149.5:3306/awa_dev
ENV FFPUPDATER_DATABASE_URI=flyawa:Nf@2el3w@tcp(35.176.149.5:3306)/awa_dev
ENV SRC_DIR=/go/src/github.com/hackstock/ffp-updater-service
ADD . ${SRC_DIR}
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR ${SRC_DIR}
RUN dep ensure -v
RUN go build -race .
ENTRYPOINT [ "./ffp-updater-service" ]