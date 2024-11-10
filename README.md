<div id="top"></div>

## 使用技術一覧

<!-- シールド一覧 -->
<!-- 該当するプロジェクトの中から任意のものを選ぶ-->
<p style="display: inline">
  <!-- バックエンドのフレームワーク一覧 -->
  <img src="https://img.shields.io/badge/-Echo-76E1FE.svg?logo=go&style=flat-square">
  <!-- バックエンドの言語一覧 -->
  <img src="https://img.shields.io/badge/-Go-76E1FE.svg?logo=go&style=flat-square">
  <!-- インフラ一覧 -->
  <img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=flat-square">
  <img src="https://img.shields.io/badge/-OpenAI-e5e5e5.svg?logo=openAI&style=flat-square">
  <img src="https://img.shields.io/badge/-Qdrant-4479A1.svg?logo=db&style=flat-square">
</p>

## 目次

1. [プロジェクトについて](#プロジェクトについて)
2. [環境](#環境)
3. [ディレクトリ構成](#ディレクトリ構成)
4. [開発環境構築](#開発環境構築)
5. [トラブルシューティング](#トラブルシューティング)

## プロジェクト名

rag_sample

## プロジェクトについて

Go言語とベクトルDBを使ったLLM（大規模言語モデル）のRAG（Retrieval Augmented Generation）サンプル

## 環境

<!-- 言語、フレームワーク、ミドルウェア、インフラの一覧とバージョンを記載 -->

| 言語・フレームワーク | バージョン |
| -------------------- | ---------- |
| Go                   | 1.23       |

<p align="right">(<a href="#top">トップへ</a>)</p>

## ディレクトリ構成

- レイヤーごとのテストを書きたかったので、オニオンアーキテクチャっぽくしましたが、依存性逆転の法則を意識しているわけではなく、結構適当なので、あまり参考にしないでください。
<!-- Treeコマンドを使ってディレクトリ構成を記載 -->

```
$ tree
.
├── Dockerfile
├── README.md
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── model
│   │   │   └── article.go
│   │   ├── repository
│   │   │   └── article_repository.go
│   │   └── service
│   │       └── openai_service.go
│   ├── infrastructure
│   │   ├── openai
│   │   │   ├── openai_service.go
│   │   │   └── openai_service_test.go
│   │   ├── qdrant
│   │   │   └── client.go
│   │   ├── repository
│   │   │   ├── article_repository.go
│   │   │   ├── article_repository_test.go
│   │   │   └── testdata
│   │   │       └── fixtures
│   │   │           └── article
│   │   │               └── find
│   │   │                   └── sample.json
│   │   └── webscraper
│   │       ├── scraper_service.go
│   │       └── scraper_service_test.go
│   ├── interface
│   │   └── controller
│   │       ├── article_controller.go
│   │       └── request.go
│   └── usecase
│       └── article_usecase.go
├── main.go
├── router
│   └── router.go
└── tmp
    ├── build-errors.log
    └── main

```

<p align="right">(<a href="#top">トップへ</a>)</p>

## 開発環境構築

<!-- コンテナの作成方法、パッケージのインストール方法など、開発環境構築に必要な情報を記載 -->

### コンテナの作成と起動

.env ファイルを以下の環境変数例と[環境変数の一覧](#環境変数の一覧)を元に作成

.env
OPENAI_API_KEY= [openAIのAPIキー]

.env ファイルを作成後、以下のコマンドで開発環境を構築

$ docker compose up --build

### 動作確認

terminalで

```
$ curl http://localhost:6333/
```

でレスポンスが返ってきたら成功

### 使い方

### コンテナの停止

以下のコマンドでコンテナを停止することができます

```
$ docker compose up
```

で起動したコンテナが動いている場合は、macの場合は[command + c ]で停止できます
