# golang 言語について思うこと

https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/02.2.html
を見ながらgolangのキャッチアップを試みている。思ったことを雑に書き留めておく。後でgolangに詳しくなった時に見返してこれらの印象が正しいものであったかを確認するのが面白いかもしれない。

## 言語仕様
- var は := による変数定義があるため使う場面はないと思う。ただし関数内出なければコンパイルが通らない
    変数を宣言する際に必要そう。:= 形式では型の宣言ができないが var 形式だと `var i int = 1` みたいにできる。

- ' " ` の違い
 	- ' ではrune(int32)型として定義される。文字コードをそのまま扱いたいときに使える？使い道そんなにないと思うけど、、、
 	- " では普通のstring。エンコードは必ずutf-8で行われるようである。必須。
 	- ` では raw string になる。改行コードなどをそのまま出力することができる。使えそう。

- iota とか言って、enum宣言時にイテレータの弱体化版みたいなのが用意されているけどいらなそう。

- 言語仕様として配列スライスが使えるのは素晴らしい。

- try catch は存在せずに exception を明示的に関数の戻り値に指定する方法をとる。panic や recover を用いて try catch を表現することもできるらしいがメジャーではないみたい。正直これは好きじゃない。

- struct 型に対して
    ```golang
    type Circle struct {
        radius float64
    }
    func (c Circle) area() float64 {
        return math.Pi * c.radius * c.radius
    }
    ```
    とかすると `c.area()` って感じで関数が呼べる。これは特徴的だと思う。ふつうはコンストラクタに関数を定義するけど、golang はそもそも後から定義するのがデフォルト？クラス記法はまた別なのかはわからん。

- type
    ```golang
    type months map[string]int
    m := months {
        "January":31,
        "February":28,
        ...
        "December":31,
    }
    func (m months) p() {
        fmt.Println(m)
    }
    ```
    こんな感じで型の定義と関数の定義もできるらしい。こんなことができるなら golang の type は実質変数として扱われてる気がする。

- go のポインタについて。go では関数内で引数の値を変更するためにはポインタ型を受け取る必要がある。が、別に変更の際にポインタ演算子を使う必要はないらしい。一種の構文糖。
    ```golang
    ```

- cとかでも同じなんだけど、アドレス演算子である `&` はどの変数に対しても定義できるみたい。構文解析的に `&&&x` でアドレスをとるとかはできないけど、
    ```golang
    x := 3
    ax := &x
    aax := &ax
    aaax := &aax
    ```
    みたいなことができ、この時 `***aaax == **aax == *ax == x` になる。使い道のなさそうな知識。

- golang では並行処理がめちゃくちゃ重要な感じで、その中でも goroutine と channel が重要。channel 変数は `chan<-` か `<-chan` で定義できるけど、ここ混乱ポイント。この変数を主体としてそれぞれ受信専用、送信専用の channel を作成、って感じ。

- golang の channel 操作のアンチパターン https://hori-ryota.com/blog/golang-channel-pattern/


