# pr-go 

NavCom analytics client for Go.

## Installation

The package can be simply installed via go get, we recommend that you use a
package version management system like the Go vendor directory or a tool like
Godep to avoid issues related to API breaking changes introduced between major
versions of the library.

To install it in the GOPATH:
```
go get https://github.com/fuziontech/pr-go
```

## Documentation

So much work to be done here 

## Usage

```go
package main

import (
    "os"

    "github.com/fuziontech/pr-go"
)

func main() {
    // Instantiates a client to use send messages to the navcom API.
    client := analytics.New(os.Getenv("NAVCOM_DSN"))

    // Enqueues a track event that will be sent asynchronously.
    client.Enqueue(analytics.Track{
        UserId: "test-user",
        Event:  "test-snippet",
    })

    // Flushes any queued messages and closes the client.
    client.Close()
}
```

## License

The library is released under the [MIT license](License.md).
