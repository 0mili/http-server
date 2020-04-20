<h1 align="center">mili Bot - HTTP Server</h1>
<p align="center">Providing HTTP integrations for mili. https://github.com/0mili/mili</p>
<p align="center">
	<a href="https://github.com/0mili/http-server/releases"><img src="https://img.shields.io/github/tag/0mili/http-server.svg?label=version&color=brightgreen"></a>
	<a href="https://circleci.com/gh/0mili/http-server/tree/master"><img src="https://circleci.com/gh/0mili/http-server/tree/master.svg?style=shield"></a>
	<a href="https://goreportcard.com/report/github.com/0mili/http-server"><img src="https://goreportcard.com/badge/github.com/0mili/http-server"></a>
    <a href="https://codecov.io/gh/0mili/http-server"><img src="https://codecov.io/gh/0mili/http-server/branch/master/graph/badge.svg"/></a>
	<a href="https://pkg.go.dev/github.com/0mili/http-server?tab=doc"><img src="https://img.shields.io/badge/godoc-reference-blue.svg?color=blue"></a>
	<a href="https://github.com/0mili/http-server/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-BSD--3--Clause-blue.svg"></a>
</p>

---

This repository contains a module for the [mili Bot library][mili].

## Getting Started

This library is packaged as [Go module][go-modules]. You can get it via:

```bash
go get github.com/0mili/http-server
```

### Example usage

In order to let your bot listen to HTTP requests you should pass the `http.Server(…)`
module when creating a new bot:

```go
package main

import (
	"github.com/0mili/mili"
	"github.com/0mili/http-server"
)

func main() {
	b := mili.New("example-bot",
		milihttp.Server("localhost:12345"),
		…
	)
	
	err := b.Run()
	if err != nil {
		b.Logger.Fatal(err.Error())
	}
}
```

When the server receives a request, it will emit it to the bots brain as `milihttp.RequestEvent`.

## Built With

* [zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging in Go
* [testify](https://github.com/stretchr/testify) - A simple unit test library

## Contributing

If you want to hack on this repository, please read the short [CONTRIBUTING.md](CONTRIBUTING.md)
guide first.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available,
see the [tags on this repository][tags]. 

## Authors

- **Friedrich Große** - *Initial work* - [fgrosse](https://github.com/fgrosse)

See also the list of [contributors][contributors] who participated in this project.

## License

This project is licensed under the BSD-3-Clause License - see the [LICENSE](LICENSE) file for details.

[mili]: https://github.com/0mili/mili
[go-modules]: https://github.com/golang/go/wiki/Modules
[tags]: https://github.com/0mili/http-server/tags
[contributors]: https://github.com/0mili/http-server/contributors
