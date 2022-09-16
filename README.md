# Translo Translation API
Go Client Library for Translo API. https://rapidapi.com/armangokka/api/translo

# Installing

```sh
go get -u github.com/transloapi/go-translo
```

# Usage

```go
package main

import (
  "github.com/transloapi/go-translo"
  "context"
  "fmt"
)

func main() {
    ctx := context.Background()
    translator := translo.NewAPI("YOUR-RAPIDAPI-KEY")
    fmt.Println(translator.Translate(ctx, "auto", "en", "Привет")) // "Hello"
    fmt.Println(translator.Detect(ctx, "Some text")) // "en"
    fmt.Println(translator.BatchTranslate(ctx, []translo.Batch{
        {
            From: "auto",
            To:   "es",
            Text: "banana",
        },
        {
            From: "auto",
            To:   "en",
            Text: "こんにちは。元気ですか？",
        },
    })) // []translo.Batch{...}
}
```
