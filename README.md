# barcodefyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/barcodefyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/barcodefyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [BarcodeFYI](https://barcodefyi.com) REST API. Barcode formats, EAN, UPC, ISBN. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [barcodefyi.com](https://barcodefyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/barcodefyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    barcodefyi "github.com/fyipedia/barcodefyi-go"
)

func main() {
    client := barcodefyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Algorithms()` | List algorithms |
| `Categories()` | List categories |
| `Components()` | List components |
| `Families()` | List families |
| `Faqs()` | List faqs |
| `Glossary()` | List glossary |
| `Gs1Prefixes()` | List gs1 prefixes |
| `Guides()` | List guides |
| `Industries()` | List industries |
| `Standards()` | List standards |
| `Symbologies()` | List symbologies |
| `Tools()` | List tools |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install barcodefyi` | [PyPI](https://pypi.org/project/barcodefyi/) |
| **npm** | `npm install barcodefyi` | [npm](https://www.npmjs.com/package/barcodefyi) |
| **Go** | `go get github.com/fyipedia/barcodefyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/barcodefyi-go) |
| **Rust** | `cargo add barcodefyi` | [crates.io](https://crates.io/crates/barcodefyi) |
| **Ruby** | `gem install barcodefyi` | [rubygems](https://rubygems.org/gems/barcodefyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| **BarcodeFYI** | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## License

MIT
