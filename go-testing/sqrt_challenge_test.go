package sqrt

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
}
func TestMany(t *testing.T) {

	records, err := GetTestCases()
	require.NoError(t, err)
	for _, tc := range records {

		value, err1 := strconv.ParseFloat(strings.TrimSpace(tc[0]), 64)
		expected, err2 := strconv.ParseFloat(strings.TrimSpace(tc[1]), 64)
		require.NoError(t, err1)
		require.NoError(t, err2)

		t.Run(fmt.Sprintf("%f", value), func(t *testing.T) {
			out, err := Sqrt(value)
			require.NoError(t, err)
			require.InDelta(t, expected, out, 0.001)

		})
	}
}

func GetTestCases() ([][]string, error) {
	data, err := ioutil.ReadFile("./testcases.csv")
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(strings.NewReader(string(data)))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
