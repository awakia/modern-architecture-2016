# Nginx API Router

## 基本設計

以下のように振り分ける

| サーバーの言語 | デフォルトのURL | Nginx経由のURL |
|--|--|--|
| rails-api | http://localhost:3000 | http://localhost/api/rails |
| go-api | http://localhost:8000 | http://localhost/api/go |
