package atc

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const url = "https://www.zva.gov.lv/zvais/zr-atc/api/atc-codes-zp/?v=csv&q="

// Entry represents the Anatomical Therapeutic Chemical as published by Zāļu Valsts Aģentūra (ZVA) in Latvia.
type Entry struct {
	Code    string
	NameEng string // name in English
	NameLat string // name in Latvian
	Level   int    // tree depth level
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

// transform transforms downloaded CSV file into []ATC.
func transform(rc io.ReadCloser) ([]Entry, error) {
	r := csv.NewReader(rc)
	r.Comma = ';'

	var (
		atc                           []Entry
		code, nameEng, nameLat, level int
		isColumnRow                   bool = true
	)

	for {
		row, err := r.Read()

		switch {
		default:
			l, err := strconv.Atoi(row[level])
			if err != nil {
				return nil, fmt.Errorf("convert level to int: %v", err)
			}

			atc = append(atc, Entry{
				Code:    row[code],
				NameLat: row[nameLat],
				NameEng: row[nameEng],
				Level:   l,
			})
		case err == io.EOF:
			rc.Close()
			return atc, nil
		case err != nil:
			return nil, fmt.Errorf("transform atc code: %w", err)
		case isColumnRow:
			for i, v := range row {
				switch v {
				case "code":
					code = i
				case "name_eng":
					nameEng = i
				case "name_lat":
					nameLat = i
				case "level":
					level = i
				}
			}
			isColumnRow = false
		}
	}
}

// Get returns []ATC - it downloads CSV file from ZVA,
// and transforms downloaded CSV into []ATC.
func Get(ctx context.Context) ([]Entry, error) {
	rc, err := extract(ctx)
	if err != nil {
		return nil, err
	}

	return transform(rc)
}
