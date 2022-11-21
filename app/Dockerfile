FROM golang:1.19-alpine

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /var/www/app

COPY . .

# ModuleモードをON
ENV GO111MODULE=on

RUN apk add build-base

EXPOSE 8080

# Airをインストールし、コンテナ起動時に実行する
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]