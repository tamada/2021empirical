# randomizer

次の問題の回答プログラム．

Go ならわかるシステムプログラミングに掲載されている次の問題の回答プログラム．

### p.74 Q4.1 タイマー

> time パッケージの time.After(duration) により、指定した時間後に時刻データを流すチャネルが得られます。
> 引数の duration は時間間隔を表す time.Duration 型で、10 * time.Second で 10 秒になります。
> time.After(duration)を使って、決まった時間を計るタイマーを作ってみま しょう。

## Usage

```sh
nml [OPTIONS] <NUMBER> [-- <COMMANDS...>]
nml means 'notify me, later!'
OPTIONS
    -u, --unit <UNIT>     specifies the time unit. Default is min
                          Available units are: msec, sec, min, and hour.
    -H, --with-header     shows the header on notification.
    -h, --help            prints this message.
NUMBER
    specifies the number for timer.
COMMANDS
    specifies the commands execute after timer.
    If no commands are specified, nml notifies by printing message
    "MESSAGE FROM ARRIVE!! (<NUMBER> <UNIT> left)" to STDOUT.
```

## Compile

```
make
```
