# json::REST::server

## 概要

* オンメモリで動くKVSです
* Json-Textのみ収納できます
* HTTPプロトコルで登録・変更・削除が出来ます
* REST APIです


## 使い方

### 起動方法

```
go run json_rest_server.go

```


### 起動オプション
-port をつけると、サーバーの待ち受けポートが変更できます。デフォルトはPort:8080/tcpです。

```
go run json_rest_server.go -port 8888
```


### データ登録
* POSTパラメータで登録します。
* pathがKeyになり、リクエストボディが登録されるjson(Value)になります。
* 同一KeyでPOSTすると、上書きされます。
* 不正なJsonをPOSTすると、エラーになります。
* 二階層以上のpath名は登録に失敗します。
* path名に「.」をつけるのは問題ありません。


### データ閲覧
* GETパラメータでデータ閲覧できます。
* pathがKeyになります。
* 存在しないKeyを指定するとエラーになります。
* 「/debug」を渡すと、登録されているKey:Valueが全出力されます。(サーバーコンソール上)


### データ削除
* DELETEパラメータでデータ削除できます。
* 存在しないKeyを削除しても問題ありません。(何も変化しないだけです)


## 検証方法
```
curl -X POST http://localhost:8080/hogehoge/poe -d '{"test":"hogehoge"}'
curl -X POST http://localhost:8080/hogehoge -d '{"test":"hogehoge"}'
curl -X POST http://localhost:8080/hogehoge2 -d '{"test2":"fuga"}'

curl http://localhost:8080/debug
curl -X DELETE http://localhost:8080/hogehoge
curl -X DELETE http://localhost:8080/hogehoge2
```

## ToDo

* messagepackとか対応してみようかな
* 管理コマンドももう少し充実

