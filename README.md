# xcd-go
eXplore and Change Directory - rewritten in Go

## 概要
ディレクトリの確認と移動が同時に行えるCLIツールです。

元々シェルスクリプト製だった[xcd](https://github.com/saka-naname/xcd)をGoで書き直してスムーズに動作するように改善しました。

## インストール
> [!IMPORTANT]
> Go 1.22+ が必要です

```bash
go install github.com/saka-naname/xcd-go@latest
```
上記のコマンドを実行後、使用しているシェルに応じて次の設定を行ってください。\
なお、使用しているGoの環境に応じて `$GOPATH/bin` を `$GOROOT/bin` や `$GOBIN` などに置き換えてください。

### Bash
`~/.bashrc` に以下の文を追記してください。
```sh
function xcd() {
    cd $($GOPATH/bin/xcd-go $@)
}
```

### zsh
`~/.zshrc` に以下の文を追記してください。
```sh
function xcd() {
    cd $($GOPATH/bin/xcd-go $@)
}
```

## 操作方法
- `↑` `↓` : 項目の選択
- `→` : 選択中のディレクトリを開く 
- `←` : 1階層上のディレクトリへ
- `Enter` : 移動先のディレクトリを確定 (終了)
- `q` : キャンセルして終了

## 今後やりたいこと
- [ ] 画面下部に操作や各種情報を表示したい
- [ ] オプションからパスを指定して好きなディレクトリから探索できるように
- [ ] 見た目をもう少しリッチにしたい
- [ ] シンボリックリンクへの対応
- [ ] 検索(フィルタ)機能の実装
