ch2
===

TCPスキャンを実施する
nmapが公開している(scanme.nmap.org)[scanme.nmap.org]があるのでスキャン対象として使う.

まずは通信できるか確認するためのプログラムを書く。

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		panic(nil)
	}
	fmt.Println("Coonnection successful")
}
```

---

次にポートスキャン。
ここはテストとして1024ポートまでとする.

```go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.DialTimeout("tcp", address, 10*time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

```

---

次はgoroutineを使った早いバージョン


```go
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.DialTimeout("tcp", address, 10*time.Second)
			if err != nil {
				panic(err)
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}

```

しかし早すぎてアウトプットが間に合わない.

そこで修理したのが次のものだ

wg(waitgroup)を利用している

```go
package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}

```

---

次にワーカーモデルを利用した物を紹介したい
```go
package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

```

## プロキシを作る
まずは入出力をGo言語で操作するために、標準入力から受け取った文字列を標準出力に出すコードを書く.

```go
package main

import (
	"fmt"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	input := make([]byte, 4096)

	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Write %d bytes to stdout\n", s)
}

```