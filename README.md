# go-markdown-emoji

[![Tag](https://img.shields.io/github/tag/4strodev/go-markdown-emoji.svg)](https://github.com/4strodev/go-markdown-emoji/releases/)
[![License](https://img.shields.io/github/license/4strodev/go-markdown-emoji.svg)](LICENSE)
[![codecov.io](https://img.shields.io/codecov/c/github/4strodev/go-markdown-emoji.svg)](https://codecov.io/github/4strodev/go-markdown-emoji)

Go module to add emoji support to [go-markdown](https://github.com/gomarkdown/markdown).


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-markdown-emoji` is a standard Go module which can be installed with:

```sh
go get github.com/4strodev/go-markdown-emoji
```

## Usage

`go-markdown-emoji` provides parser and renderer hooks to the markdown engine.  The parser hook is `Parser` and the renderer `Renderer`.

Emojis are signified in Markdown as names between colons, for example `:smile:`.  A full list of the emojis supported can be seen in [emoji.go](https://github.com/4strodev/go-markdown-emoji/blob/master/emoji.go).

### Example

```go
package main

import (
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/parser"
    emoji "github.com/4strodev/go-markdown-emoji"
)

func main() {
    p := parser.New()
    p.Opts = parser.Options{ParserHook: emoji.Parser}
    r := html.NewRenderer(html.RendererOptions{
      Flags: html.CommonFlags,
      RenderNodeHook: emoji.Renderer,
    })

    html := markdown.ToHTML([]byte(":smile:"), p, r)
    fmt.Printf("%s\n", string(html))
}
```

## Maintainers
- (Author) Jim McDonald: [@mcdee](https://github.com/mcdee).
- Alejandro Marin: [@4strodev](https://github.com/4strodev).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/4strodev/go-markdown-emoji/issues).

## License

[Apache-2.0](LICENSE).
