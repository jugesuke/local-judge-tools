# local-judge-tools

## Feature
ローカルで動く、競プロライクなジャッジシステムです。

## Setup

### マシンに直接インストールする場合

1. このリポジトリをCloneする
1. `testcases/data`にテストケースを入れる
1. installする

    通常版を利用する場合は、
    ```
    go install ./cmd/lj
    ```
    tiny版を利用する場合は、
    ```
    go install ./cmd/tiny-lj
    ```

### 実行ファイルを自分で設置する場合
1. このリポジトリをCloneする
1. `testcases/data`にテストケースを入れる
1. ビルドする

    通常版を利用する場合は、
    ```
    go build ./cmd/lj
    ```
    tiny版を利用する場合は、
    ```
    go build ./cmd/tiny-lj
    ```
1. お好きな、pathの通ったディレクトリに配置する
