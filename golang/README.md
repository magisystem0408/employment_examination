# 採用試験【Go】

TRUNKに興味を持っていただいて有り難うございます。
本テストは、Go言語に関する基礎的な知識、Webの基礎知識、単体テストの知識があることを判断することが目的です。
問題は２つで、`In-memory DB`と`APIサーバ`に関して作成して頂きますが、実装の中身は自由になっています。
単体テストも書いて頂きますようお願いいたします。


## 問題

セットアップと提出は、コードが書かれているので参考にしてください。
問題文は、「`In-memory DB` の作成」と「`APIサーバ` の作成」です。

### セットアップ

1. リポジトリをフォークして、ブランチ作成
```bash
# https://github.com/trunk-inc/employment_examinationをフォークする。
git clone https://github.com/<GitHubのアカウント名>/employment_examination
cd employment_examination/golang
git checkout -b exam-golang
```

2. Goモジュールの初期化
```bash
go mod init
```

### `In-memory DB` の作成

以下のGetとSetのインターフェイスを持った簡易なKVSを作成してください。
```go
type DB interface {
	Get(key string) (value string, err error)
	Set(key, value string) error
}
```

ファイル名は`db.go`、テストファイルは`db_test.go`がおすすめです。

### `APIサーバ` の作成

以下の二点を満たすサーバを実装してください。
1. `/healthz`にアクセスすると、ステータスコード200の`ok`が返ってくる。
2. 先ほど作成したIn-memoryDBを利用し、ユーザの登録とログインができるAPIを作成
* `[POST] /users/register` で登録し、`ok`を返す。
* `[POST] /users/login` でログインし、ユーザ情報を返す。
* Bodyには、メールアドレスとパスワードを含む
* パスワードはハッシュ化する

### 提出

```bash
git push -u origin exam-golang
open https://github.com/trunk-inc/employment_examination/pulls
# この後、プルリクエストを作成してください。
```
