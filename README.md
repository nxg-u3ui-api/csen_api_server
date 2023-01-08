# apibridge

web hook で通知されるイベントを、websocketでつながったクライアントに同報周知する。

web hookの通知は、自身がRESTサーバーとして受付け、json形式のデータがのっていることを想定

## REST APIサーバー（web hook client）

http://host:port/api
で受け付ける。

## websocket 接続

http://host:port/ws
で受け付ける。

## 動作確認

### 前条件

* golang (1.13.8）
* wscatがインストールされている

### 本サーバーアプリの起動

./run.sh

### websocketクライアントの起動&接続

./client.sh

### web hook での情報Push

./pusher.sh
