# qiita-body-replace
自分が投稿した全てのQiita/QiitaTeam記事の本文中の文字列を一括置換します

## 使い方
### 1. Qiitaアクセストークンを取得する
https://qiita.com/settings/tokens/new にアクセスして  
アクセストークン名を設定し、更新したい対象(Qiita or Qiita:Team)のread/writeスコープにチェックを入れる。  

<img width="500" alt="mosaic_token" src="https://user-images.githubusercontent.com/4005383/51170130-01ecec80-18f1-11e9-9b95-b6b3ce807e19.png">

表示されたトークンを控える。  

### 2. いずれかの方法で取得する
#### a. `go get`で導入する
```
$ go get -u github.com/uskey512/qiita-body-replace
$ qiita-body-replace
```

#### b. ビルド版を入手する
- [qiita-body-replace.darwin.386 5.47 MB](https://github.com/uskey512/qiita-body-replace/releases/download/v1.0.0/qiita-body-replace.darwin.386)
- [qiita-body-replace.darwin.amd64 6.32 MB](https://github.com/uskey512/qiita-body-replace/releases/download/v1.0.0/qiita-body-replace.darwin.amd64)
- [qiita-body-replace.windows.386.exe 5.24 MB](https://github.com/uskey512/qiita-body-replace/releases/download/v1.0.0/qiita-body-replace.windows.386.exe)
- [qiita-body-replace.windows.amd64.exe 6.03 MB](https://github.com/uskey512/qiita-body-replace/releases/download/v1.0.0/qiita-body-replace.windows.amd64.exe)


### 3. 実行する
下記のパラメータを入力すると置換が開始されます。  
- アクセストークン
- 置換前文字列
- 置換後文字列
- 所属チーム(Qiita:Teamの場合のみ)
  - `example`.qiita.com の `example`部分

終了時に置換された記事数 / 全記事数が表示されます。  

### 利用例
#### 参照している外部サイトが別ドメインに移行してしまった場合
`http://aaa.com/`から`https://bbb.com/`に移行したケースだと  
変更前文字列 : `http://aaa.com/`  
変更後文字列 : `http://bbb.com/`  
として実行します。  
