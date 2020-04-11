# Exploiting DNS

DNS周りの操作をGo言語でやってみる

## Aレコードを確認するコードを書いてみる

`github.com/miekg/dns`パッケージを使ってAレコードを確認するコードがコチラ

```go
package main

import (
	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")

	msg.SetQuestion(fqdn, dns.TypeA)
	dns.Exchange(&msg, "8.8.8.8:53")
}
```

下記内容でtcp dumpとっていればわかる

```sh
$ sudo tcpdump -i enp0s31f6 -n udp port 53
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
---snip---
16:46:18.146627 IP 192.168.1.8.60704 > 8.8.8.8.53: 6428+ A? stacktitan.com. (32)
16:46:18.195156 IP 8.8.8.8.53 > 192.168.1.8.60704: 6428 1/0/0 A 34.212.50.84 (48)
---snip
```

## Aレコードを出力してみる

```go
package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg

	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}

	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a.A)
		}
	}
}
```