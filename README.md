# Anatomical Therapeutic Chemical (ATC) Classification System [<img src="https://github.com/google/region-flags/blob/gh-pages/svg/LV.svg" width="45" height="30" alt="Latviski">](https://github.com/expect-digital/atc/README.lv.md)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/expect-digital/atc/Check) ![GitHub](https://img.shields.io/github/license/expect-digital/atc)

This ```atc``` package provides easy and fast way to extract and transform (ET from [ETL](https://en.wikipedia.org/wiki/Extract,_transform,_load)) ATC data into ```golang``` struct. [The ATC list](https://www.zva.gov.lv/lv/veselibas-aprupes-specialistiem-un-iestadem/zales/atk-klasifikacija) is maintained and published by [ZVA](https://www.zva.gov.lv/en) in Latvia.

For example:

```go
package main

import (
  "fmt"

  "github.com/expect-digital/atc"
)

func main() {
  entries, err := atc.Get()
  if err != nil {
    panic(err)
  }

  fmt.Printf("%+v", entries)
}
```