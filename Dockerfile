# Goビルド用のベースイメージを指定
FROM golang:1.20 AS build

# 作業ディレクトリを設定
WORKDIR /app

# Goの依存関係をコピー
COPY go.mod ./
RUN go mod download

# ソースコードをコピー
COPY . .

