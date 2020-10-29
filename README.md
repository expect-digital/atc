# atc ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/expect-digital/atc/Check) ![GitHub](https://img.shields.io/github/license/expect-digital/atc)

 ```atc``` package provides easy and fast way to extract and transform (ET from [ETL](https://en.wikipedia.org/wiki/Extract,_transform,_load)) Anatomical Therapeutic Chemical (ATC) Classification System data into ```golang``` values. [The ATC list](https://www.zva.gov.lv/lv/veselibas-aprupes-specialistiem-un-iestadem/zales/atk-klasifikacija) is maintained and published by [ZVA](https://www.zva.gov.lv/en) in Latvia.

For example:

```go
entries, err := atc.GetEntries(context.Background())
if err != nil {
  panic(err)
}

fmt.Printf("%+v\n", entries)
```

Use ```atc.Get()``` to map values into your own struct - ```csv``` field tags as defined by [csvutil](https://github.com/jszwec/csvutil). Please look at ```atc.Entry``` for all available values that can be extracted from ATC.

```go
var entries []struct{
  Code `csv:"code"`
  Name `csv:"name_eng"`
}

err := atc.Get(context.Background(), &entries)
if err != nil {
  panic(err)
}
fmt.Printf("%+v\n", entries)
```