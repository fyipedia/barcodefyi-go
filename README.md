# barcodefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/barcodefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/barcodefyi-go)

Go client for the [BarcodeFYI](https://barcodefyi.com) API. Look up barcode symbologies, standards, components, and industry applications. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/barcodefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    barcodefyi "github.com/fyipedia/barcodefyi-go"
)

func main() {
    client := barcodefyi.NewClient()

    result, err := client.Search("upc")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'upc'\n", result.Total)

    sym, err := client.Symbology("upc-a")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", sym.Name, sym.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search symbologies, standards, and glossary |
| `Symbology(slug)` | Get barcode symbology details |
| `Family(slug)` | Get barcode family details |
| `Standard(slug)` | Get barcode standard details |
| `Component(slug)` | Get barcode component details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two symbologies |
| `Random()` | Get a random symbology |
| `Industry(slug)` | Get industry application details |

## REST API

Free, no authentication required. Base URL: `https://barcodefyi.com/api`

```bash
curl https://barcodefyi.com/api/search/?q=upc
curl https://barcodefyi.com/api/symbology/upc-a/
curl https://barcodefyi.com/api/random/
```

Full OpenAPI spec: https://barcodefyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [barcodefyi](https://pypi.org/project/barcodefyi/) | `pip install barcodefyi` |
| TypeScript | [barcodefyi](https://www.npmjs.com/package/barcodefyi) | `npm install barcodefyi` |
| Go | [barcodefyi-go](https://pkg.go.dev/github.com/fyipedia/barcodefyi-go) | `go get github.com/fyipedia/barcodefyi-go` |
| Rust | [barcodefyi](https://crates.io/crates/barcodefyi) | `cargo add barcodefyi` |
| Ruby | [barcodefyi](https://rubygems.org/gems/barcodefyi) | `gem install barcodefyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
