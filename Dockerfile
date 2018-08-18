FROM golang:1.10.2-alpine

RUN apk add --no-cache chromium chromium-chromedriver git bash && \
    go get -u github.com/golang/dep/cmd/dep && \
    go get -u github.com/golang/lint/golint

WORKDIR /go/src/check_medical_checkup_booking
COPY . /go/src/check_medical_checkup_booking
