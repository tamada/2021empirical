# copier

Goならわかるシステムプログラミングに掲載されている次の問題の回答プログラム．

### p.59 Q3.1 ファイルのコピー

> 古いファイル(old.txt)を新しいファイル(new.txt)にコピーしてみましょう。 本章で紹介したサンプルコードを応用すれば難しくないと思います。
> さらに改造して実用的なコマンドにしてみたいと思われる方は、コマンドラインオプションでファイル名を渡せるようにするとよいでしょう。本書の範囲からは外れるので詳細は省きますが、os.Args という文字列配列にオプションが格納されます。また、標準ライブラリにある flag パッケージを使うと、オプションのパース処理がより便利に行えます。

## Usage

```sh
copier <FROM> <TO>
```

## Compile

```
make
```
