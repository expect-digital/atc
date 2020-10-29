package atc

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transform(t *testing.T) {
	var entries []Entry
	err := transform(strings.NewReader(`code;name_eng;name_lat;level
A;"ALIMENTARY TRACT AND METABOLISM";"Gremošanas trakts un vielmaiņa";1
A01;"STOMATOLOGICAL PREPARATIONS";"Stomatoloģiskie līdzekļi";2`), &entries)

	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		{"A", "ALIMENTARY TRACT AND METABOLISM", "Gremošanas trakts un vielmaiņa", 1},
		{"A01", "STOMATOLOGICAL PREPARATIONS", "Stomatoloģiskie līdzekļi", 2}}, entries)
}

func Test_Get(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping")
		return
	}

	var entries []struct {
		C string `csv:"code"`
		N string `csv:"name_eng"`
	}

	assert.NoError(t, Get(context.Background(), &entries))
	assert.NotEmpty(t, entries)

	for _, v := range entries {
		assert.NotEmpty(t, v.C)
		assert.NotEmpty(t, v.N)
	}
}

func Test_GetEntries(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping")
		return
	}

	entries, err := GetEntries(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, entries)

	// not all names have been translated into Latvian
	var nameLatFound bool

	for _, v := range entries {
		assert.NotEmpty(t, v.Code)
		assert.NotEmpty(t, v.NameEng)
		if len(v.NameLat) > 0 {
			nameLatFound = true
		}
		assert.NotEmpty(t, v.NameEng)
	}

	assert.True(t, nameLatFound)
}
