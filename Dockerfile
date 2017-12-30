FROM golang:onbuild

ADD main.go index.html /go/src/github.com/yvrdev/go-slack-invite/
RUN go install github.com/yvrdev/go-slack-invite
ENTRYPOINT /go/bin/go-slack-invite

EXPOSE 3000

