package atc

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"github.com/dimchansky/utfbom"
	"github.com/jszwec/csvutil"
)

const url = "https://www.zva.gov.lv/zvais/zr-atc/api/atc-codes-zp/?v=csv&q="

// Entry represents the Anatomical Therapeutic Chemical as published by Zāļu Valsts Aģentūra (ZVA) in Latvia.
type Entry struct {
	Code    string `csv:"code"`
	NameEng string `csv:"name_eng"` // name in English
	NameLat string `csv:"name_lat"` // name in Latvian, not all names have been translated into Latvian
	Level   int    `csv:"level"`    // tree depth level
}

// extract downloads CSV file from ZVA.
func extract(ctx context.Context) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create ATC csv request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get ATC csv: %w", err)
	}

	return res.Body, nil
}

// bomReader helps to remove double BOM from CSV file. utfbom does
// NOT apply second Skip() if it is already *utfbom.Reader type.
type bomReader struct {
	r io.Reader
}

// Read implements io.Reader interface.
func (b *bomReader) Read(p []byte) (n int, err error) {
	return b.r.Read(p)
}

// transform transforms downloaded CSV file into dest slice.
func transform(rc io.Reader, dest interface{}) error {
	// skip double BOM in CSV file
	r := csv.NewReader(utfbom.SkipOnly(&bomReader{r: utfbom.SkipOnly(rc)}))
	r.Comma = ';'

	d, err := csvutil.NewDecoder(r)
	if err != nil {
		return fmt.Errorf("create new ATC decoder: %w", err)
	}

	if err = d.Decode(&dest); err != nil {
		return fmt.Errorf("decode ATC: %w", err)
	}

	return nil
}

// Get returns ATC values mapped to given dest.
//
// Values are mapped using csv tags https://github.com/jszwec/csvutil , please
// check Entry struct for available csv column names.
//
// Every time Get is called, the CSV file is downloaded from ZVA.
func Get(ctx context.Context, dest interface{}) error {
	rc, err := extract(ctx)
	if err != nil {
		return err
	}

	defer rc.Close()

	return transform(rc, &dest)
}

// Get returns []Entry.
//
// Every time GetEntries is called, the CSV file is downloaded from ZVA.
func GetEntries(ctx context.Context) ([]Entry, error) {
	var r []Entry
	err := Get(ctx, &r)
	return r, err
}
