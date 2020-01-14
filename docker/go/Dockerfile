FROM golang:1.13.1-alpine
# GO Modulesを利用するため
ENV GO111MODULE=on
# ホットリロード環境作成のためfreshをgo get
RUN apk update && \
    apk add emacs curl git && \
    go get github.com/pilu/fresh
