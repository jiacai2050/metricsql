package metricsql

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicode"
)

const (
	HeadOp     = "op"
	HeadFunc   = "func"
	HeadAggr   = "aggr"
	HeadRollup = "rollup"
	HeadMetric = "metric"
	HeadString = "string"
	HeadNumber = "number"
)

func (e BinaryOpExpr) MarshalJSON() ([]byte, error) {
	type wrap BinaryOpExpr
	return wrapType(wrap(e), HeadOp)
}

func (e FuncExpr) MarshalJSON() ([]byte, error) {
	type wrap FuncExpr
	return wrapType(wrap(e), HeadFunc)
}

func (e AggrFuncExpr) MarshalJSON() ([]byte, error) {
	type wrap AggrFuncExpr
	return wrapType(wrap(e), HeadAggr)
}

func (e RollupExpr) MarshalJSON() ([]byte, error) {
	type wrap RollupExpr
	return wrapType(wrap(e), HeadRollup)
}

func (e MetricExpr) MarshalJSON() ([]byte, error) {
	type wrap MetricExpr
	return wrapType(wrap(e), HeadMetric)
}

func (e StringExpr) MarshalJSON() ([]byte, error) {
	type wrap StringExpr
	return wrapType(wrap(e), HeadString)
}

func (e NumberExpr) MarshalJSON() ([]byte, error) {
	type wrap NumberExpr
	return wrapType(wrap(e), HeadNumber)
}

func wrapType(v interface{}, typo string) ([]byte, error) {
	var buf bytes.Buffer

	buf.Write([]byte(fmt.Sprintf(`{"Head": "%s", "Body": `, typo)))
	bs, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	buf.Write(bs)
	buf.WriteByte('}')

	return buf.Bytes(), nil
}

func isValidUnicode(s string) bool {
	if len(s) == 0 {
		return false
	}
	return unicode.Is(unicode.Han, []rune(s)[0])
}
