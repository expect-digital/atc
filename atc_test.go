package atc

import (
	"context"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transform(t *testing.T) {
	atc, err := transform(ioutil.NopCloser(strings.NewReader(`code;name_eng;name_lat;level
A;"ALIMENTARY TRACT AND METABOLISM";"Gremošanas trakts un vielmaiņa";1
A01;"STOMATOLOGICAL PREPARATIONS";"Stomatoloģiskie līdzekļi";2`)))

	assert.NoError(t, err)
	assert.Equal(t, []ATC{
		{"A", "ALIMENTARY TRACT AND METABOLISM", "Gremošanas trakts un vielmaiņa", 1},
		{"A01", "STOMATOLOGICAL PREPARATIONS", "Stomatoloģiskie līdzekļi", 2}}, atc)
}

func Test_Get(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping")
		return
	}

	atc, err := Get(context.Background())

	assert.NoError(t, err)
	assert.NotEmpty(t, atc)
}
