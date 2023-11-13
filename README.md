# skycasino-proto

## 說明

定義Casino遊戲的紀錄以及推送到skybet的事件
實作在skycasino

## gen proto

install buf
define the flow in buf.gen.yaml

buf generate

生成 tag 需在執行 generate 後跑
buf generate --template buf.gen.tag.yaml

``` shell

pnpm gen:go
pnpm gen:gotag
pnpm prepare

```
