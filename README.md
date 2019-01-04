# Cybozu Garoon REST API client for Go [![CircleCI](https://circleci.com/gh/otoyo/garoon.svg?style=svg)](https://circleci.com/gh/otoyo/garoon)

Supported most of [Garoon REST API](https://developer.cybozu.io/hc/ja/articles/360000503306) excluding `notification/items(POST)`

```go
import "github.com/otoyo/garoon"
```

## Example

```go
package main

import (
        "fmt"

        "github.com/otoyo/garoon"
)

func main() {
        c, err := garoon.NewClient("xxx", "user", "password")
        if err != nil {
                fmt.Printf("%s\n", err)
                return
        }

        ev, err := c.FindEvent("123")
        if err != nil {
                fmt.Printf("%s\n", err)
                return
        }

        fmt.Printf("%#v\n", ev)
}
```

## Documentation

See [wiki](https://github.com/otoyo/garoon/wiki)

## Contribution

Contributions are welcome ;)
