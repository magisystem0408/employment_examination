# 採用試験【Flutter】

TRUNKに興味を持っていただいて有り難うございます。
本テストは、Flutterの基礎的な知識とGraphQLの知識を確認するものとなっています。
デザイン性は問いません。

*ただし、materialとcupertinoのウィジェットは使わないでください。*

## 問題

ユーザ登録とログイン機能、ToDoの追加とToDoの完了の４つの機能がついたWebアプリを作成

下記に流れを記載します。
セットアップと提出は、コードが書かれているので参考にしてください。

### セットアップ
1. リポジトリをフォークして、ブランチ作成
```bash
# https://github.com/trunk-inc/employment_examinationをフォークする。
git clone https://github.com/<GitHubのアカウント名>/employment_examination
cd employment_examination/flutter
git checkout -b exam-flutter
```

2. Flutterプロジェクトの追加
```bash
flutter create exam_flutter
cd exam_flutter
flutter run -d macos
```

### UIの作成
1. ユーザ名(任意のアルファベット)とパスワードの新規登録/ログイン画面
3. ToDoリスト画面

### GraphQLからDartのコード生成

まずは、サーバを起動
```
make -C ./graphql-server
```

riverpod と artemisをpub.yamlに追加
* https://pub.dev/packages/riverpod
* https://pub.dev/packages/artemis

artemisでコード生成
ユーザの状態とToDoの状態を管理するProviderを追加
4つの機能が満たされるように実装。


### 提出

```bash
git push -u origin exam-flutter
open https://github.com/trunk-inc/employment_examination/pulls
# この後、プルリクエストを作成してください。
```
