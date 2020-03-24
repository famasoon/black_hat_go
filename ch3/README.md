ch3
===

## HTTP クライアントを作ろう
GET, HEAD, PUT, DELETE, POSTができるよ

試しにGoogleの`robots.txt`を出力してみるよ

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com/robots.txt")
	if err != nil {
		log.Fatalln("Unable to accept connection")
	}

	fmt.Println("Status: " + resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body))
}

```

shodanのツールを作る
作った
shodanディレクトリを参照
(fishだとうまく動かないからbashでやること)
`shodan`ディレクトリに収まっている。

---

## Interacting with Metasploit
Metasploit は攻撃によく使われるフレームワーク。
この項目では`metasploit`についてみっちりやっていく。

Kaliの入っているLinuxマシンを使う必要があるので一旦ペンディング
* [ ] Kali linuxを立ち上げてこのトレーニングを実施する

## Parsing Document Metadata with Bing Scraping
Bing のスクレイピング結果からメタデータを抽出してみる。

https://github.com/famasoon/parseBING

に結果書いたから確認。

