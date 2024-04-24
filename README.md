# xcd-go
![GitHub Release](https://img.shields.io/github/v/release/saka-naname/xcd-go?style=for-the-badge&color=green)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/saka-naname/xcd-go?style=for-the-badge)


eXplore and Change Directory - rewritten in Go

## 概要
ディレクトリの確認と移動が同時に行えるCLIツールです。

元々シェルスクリプト製だった[xcd](https://github.com/saka-naname/xcd)をGoで書き直してスムーズに動作するように改善しました。

![20240425_050115](https://github.com/saka-naname/xcd-go/assets/61535180/93bb36e9-f6ea-452c-ad3a-5f14f37c5ace)

## インストール
> [!IMPORTANT]
> Go 1.22+ が必要です

```bash
go install github.com/saka-naname/xcd-go@latest
```
上記のコマンドを実行後、使用しているシェルに応じて次の設定を行ってください。

### Bash
`~/.bashrc` に以下の文を追記してください。
```sh
function xcd() {
    cd $($(go env GOPATH)/bin/xcd-go $@)
}
```

### zsh
`~/.zshrc` に以下の文を追記してください。
```sh
function xcd() {
    cd $($(go env GOPATH)/bin/xcd-go $@)
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
