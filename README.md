# json::REST::server

## 概要

* オンメモリで動くKVSです
* ~~Json-Textのみ収納できます~~ ← 一部のjsonが判定できないので中止
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
* POSTリクエストで登録します。
* pathがKeyになり、リクエストボディが登録されるjson(Value)になります。
* 同一KeyでPOSTすると、上書きされます。
* ~~不正なJsonをPOSTすると、エラーになります。~~
* 二階層以上のpath名は登録に失敗します。
* path名に「.」をつけるのは問題ありません。


### データ閲覧
* GETリクエストでデータ閲覧できます。
* pathがKeyになります。
* 存在しないKeyを指定するとエラーになります。
* 「/\_debug\_」を渡すと、登録されているKey:Value等が全出力されます。(サーバーコンソール上)


### データ削除
* DELETEリクエストでデータ削除できます。
* 存在しないKeyを削除しても問題ありません。(何も変化しないだけです)
* 「/\_reset\_」を渡すと、全KEYが削除されます。

## 検証方法
```
curl -X POST http://localhost:8080/hogehoge/poe -d '{"test":"hogehoge"}' <= false
curl -X POST http://localhost:8080/hogehoge -d '{"test":"hogehoge"}'
curl -X POST http://localhost:8080/hogehoge2 -d '{"test2":"fuga"}'

curl http://localhost:8080/_debug_ => DEBUG Message 
curl -X DELETE http://localhost:8080/hogehoge => {"test":"hogehoge"}
curl -X DELETE http://localhost:8080/hogehoge2 => {"test2":"fuga"}
```

## ToDo

* データ容量圧縮(gzip)
* list実装との速度比較
* messagepackとか対応
* パスワード
* Template化
