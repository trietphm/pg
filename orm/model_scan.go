package orm

import (
	"fmt"
	"reflect"

	"github.com/go-pg/pg/v10/types"
)

type scanValuesModel struct {
	Discard
	values  []interface{}
	columns map[string]types.ColumnInfo
}

var _ Model = scanValuesModel{}

//nolint
func Scan(values ...interface{}) scanValuesModel {
	return scanValuesModel{
		values: values,
	}
}

func (scanValuesModel) useQueryOne() bool {
	return true
}

func (m scanValuesModel) NextColumnScanner() ColumnScanner {
	return m
}

func (m scanValuesModel) ScanColumn(col types.ColumnInfo, rd types.Reader, n int) error {
	if int(col.Index) >= len(m.values) {
		return fmt.Errorf("pg: no Scan var for column index=%d name=%q",
			col.Index, col.Name)
	}

	m.saveColumnsInfo(col)
	return types.Scan(m.values[col.Index], rd, n)
}

func (m scanValuesModel) Columns() map[string]types.ColumnInfo {
	return m.columns
}

func (m *scanValuesModel) saveColumnsInfo(col types.ColumnInfo) {
	if m.columns == nil {
		m.columns = make(map[string]types.ColumnInfo)
	}
	if _, ok := m.columns[col.Name]; !ok {
		m.columns[col.Name] = col
	}
}

//------------------------------------------------------------------------------

type scanReflectValuesModel struct {
	Discard
	values  []reflect.Value
	columns map[string]types.ColumnInfo
}

var _ Model = scanReflectValuesModel{}

func scanReflectValues(values []reflect.Value) scanReflectValuesModel {
	return scanReflectValuesModel{
		values: values,
	}
}

func (scanReflectValuesModel) useQueryOne() bool {
	return true
}

func (m scanReflectValuesModel) NextColumnScanner() ColumnScanner {
	return m
}

func (m scanReflectValuesModel) ScanColumn(col types.ColumnInfo, rd types.Reader, n int) error {
	if int(col.Index) >= len(m.values) {
		return fmt.Errorf("pg: no Scan var for column index=%d name=%q",
			col.Index, col.Name)
	}
	return types.ScanValue(m.values[col.Index], rd, n)
}

func (m scanReflectValuesModel) Columns() map[string]types.ColumnInfo {
	return m.columns
}

func (m *scanReflectValuesModel) saveColumnsInfo(col types.ColumnInfo) {
	if m.columns == nil {
		m.columns = make(map[string]types.ColumnInfo)
	}
	if _, ok := m.columns[col.Name]; !ok {
		m.columns[col.Name] = col
	}
}
