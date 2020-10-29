# Anatomiski Terapeitiski Ķīmiskās (ATĶ) klasifikācijas sistēma [<img src="https://github.com/google/region-flags/blob/gh-pages/svg/GB.svg" width="45" height="30" alt="English">](https://github.com/expect-digital/atc/README.md)

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/expect-digital/atc/Check) ![GitHub](https://img.shields.io/github/license/expect-digital/atc)

Šī ```atc``` bibliotēka nodrošina vienkāršotu un ātru ATĶ datu lejupielādi un konvertēšanu uz ```golang``` datu struktūru. [ATĶ sarakstu](https://www.zva.gov.lv/lv/veselibas-aprupes-specialistiem-un-iestadem/zales/atk-klasifikacija) uztur un publicē [ZVA](https://www.zva.gov.lv/lv) Latvijā.

Piemēram:

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