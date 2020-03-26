まずは`hello_server`(いわゆるecho server)を作成してみる。

```go
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}

```

結果は以下の通り

```sh
$ curl -i -X GET http://localhost:8000/hello?name=test
HTTP/1.1 200 OK
Date: Wed, 25 Mar 2020 15:54:18 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8

Hello test
```

---

次にルーティングするHTTPサーバを作ってみる

```go
package main

import (
	"fmt"
	"net/http"
)

type router struct{}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/a":
		fmt.Fprintf(w, "Executed /a\n")
	case "/b":
		fmt.Fprintf(w, "Executed /b\n")
	case "/c":
		fmt.Fprintf(w, "Executed /c\n")
	default:
		http.Error(w, "404 Not Found\n", 404)
	}
}

func main() {
	var r router
	http.ListenAndServe(":8000", &r)
}

```

結果は以下の通り

```sh
$ curl -i -XGET http://localhost:8000/a
HTTP/1.1 200 OK
Date: Wed, 25 Mar 2020 16:01:17 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8

Executed /curl -i -XGET http://localhost:8000/a
HTTP/1.1 200 OK
Date: Wed, 25 Mar 2020 16:01:51 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Executed /a
$ curl -i -XGET http://localhost:8000/b
HTTP/1.1 200 OK
Date: Wed, 25 Mar 2020 16:01:55 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Executed /b
$ curl -i -XGET http://localhost:8000/c
HTTP/1.1 200 OK
Date: Wed, 25 Mar 2020 16:01:57 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Executed /c

$ curl -i -XGET http://localhost:8000/d
HTTP/1.1 404 Not Found
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 25 Mar 2020 16:01:59 GMT
Content-Length: 15

404 Not Found
```

## simple middleware
既存の構造体に`ServeHTTP`関数を追加することで`http.Handle`に登録できる。
これを利用して接続があるたびに標準出力に
`log.Println("star")`
`log.Println("finish")`
を実行させて文字列を吐き出させる。

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type logger struct {
	Inner http.Handler
}

func (l *logger) (w http.ResponseWriter, r *http.Request) {
	log.Println("star")
	l.Inner.ServeHTTP(w, r)
	log.Println("finish")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	f := http.HandlerFunc(hello)
	l := logger{Inner: f}
	http.ListenAndServe(":8080", &l)
}

```