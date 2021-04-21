# randomizer

次の問題の回答プログラム．

Goならわかるシステムプログラミングに掲載されている次の問題の回答プログラム．

### p.59 Q3.2 テスト用の適当なサイズのファイルを作成

>  ファイルを作成してランダムな内容で埋めてみましょう。
> crypto/rand パッケージ(本来は付録 A で紹介するように暗号用の機能)をインポートすると、rand.Reader という io.Reader が使えます。この Reader は、ランダムなバイトを延々と出力し続ける無限長のファイルのような動作をします。これを使って、1024 バイトの長さのバイナリファイルを作ってみましょう。
> ヒントですが、io.Copy() を使ってはいけません。io.Copy() は Reader の終了まですべて愚直にコピーしようとします。そして rand.Reader には終わりはありません。あとはわかりますよね?

## Usage

```sh
randomizer [OPTIONS]
OPTIONS
    -s, --size <SIZE>     specifies the data size for generation. Default is 1024.
    -d, --dest <DEST>     specifies the destination file. Default is stdout.
    -h, --help            prints this message.
```

## Compile

```
make
```
