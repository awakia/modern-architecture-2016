# Rails API Server

RailsでJSONを返すAPIサーバーを作る場合

## 基本設計

Rails 5 のAPI Onlyアプリケーションで作成する
http://edgeguides.rubyonrails.org/api_app.html

```
rails new server --api -TB -d postgresql
```

### モデル構造

```
rails g scaffold User name account_name:string:unique
rails g scaffold Issue title body:text author:refereces assignee:references locked:boolean comment_count:integer
rails g scaffold IssueComment body:text author:references
rails g scaffold Label name color:integer open_count:integer
```
