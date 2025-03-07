package orm

import (
	"github.com/go-pg/pg/v10/types"
)

type mapModel struct {
	hookStubs
	ptr     *map[string]interface{}
	m       map[string]interface{}
	columns map[string]types.ColumnInfo
}

var _ Model = (*mapModel)(nil)

func newMapModel(ptr *map[string]interface{}) *mapModel {
	model := &mapModel{
		ptr: ptr,
	}
	if ptr != nil {
		model.m = *ptr
	}
	return model
}

func (m *mapModel) Init() error {
	return nil
}

func (m *mapModel) NextColumnScanner() ColumnScanner {
	if m.m == nil {
		m.m = make(map[string]interface{})
		*m.ptr = m.m
	}
	return m
}

func (m mapModel) AddColumnScanner(ColumnScanner) error {
	return nil
}

func (m mapModel) Columns() map[string]types.ColumnInfo {
	return m.columns
}

func (m *mapModel) ScanColumn(col types.ColumnInfo, rd types.Reader, n int) error {
	val, err := types.ReadColumnValue(col, rd, n)
	if err != nil {
		return err
	}

	m.m[col.Name] = val
	m.saveColumnsInfo(col)
	return nil
}

func (m *mapModel) saveColumnsInfo(col types.ColumnInfo) {
	if m.columns == nil {
		m.columns = make(map[string]types.ColumnInfo)
	}
	if _, ok := m.columns[col.Name]; !ok {
		m.columns[col.Name] = col
	}
}

func (mapModel) useQueryOne() bool {
	return true
}
