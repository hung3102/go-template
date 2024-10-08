// Code generated by volcago. DO NOT EDIT.
// generated version: v1.11.0
package repositories

import (
	"time"

	"cloud.google.com/go/firestore"
	"github.com/go-utils/dedupe"
	"golang.org/x/xerrors"
	"google.golang.org/genproto/googleapis/type/latlng"
)

// QueryChainer - query chainer
type QueryChainer struct {
	QueryGroup       []*Query
	OrderByDirection firestore.Direction
	Filter           Filter
	cursor           Cursor
	err              error
}

// NewQueryChainer - constructor
func NewQueryChainer() *QueryChainer {
	return new(QueryChainer)
}

// Equal - `==`
func (qc *QueryChainer) Equal(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeEqual)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// NotEqual - `!=`
func (qc *QueryChainer) NotEqual(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeNotEqual)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// LessThan - `<`
func (qc *QueryChainer) LessThan(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeLessThan)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// LessThanOrEqual - `<=`
func (qc *QueryChainer) LessThanOrEqual(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeLessThanOrEqual)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// GreaterThan - `>`
func (qc *QueryChainer) GreaterThan(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeGreaterThan)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// GreaterThanOrEqual - `>=`
func (qc *QueryChainer) GreaterThanOrEqual(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeGreaterThanOrEqual)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// In - `in`
func (qc *QueryChainer) In(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeIn)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// NotIn - `not-in`
func (qc *QueryChainer) NotIn(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeNotIn)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// ArrayContains - `array-contains`
func (qc *QueryChainer) ArrayContains(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeArrayContains)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// ArrayContainsAny - `array-contains-any`
func (qc *QueryChainer) ArrayContainsAny(v interface{}) *QueryChainer {
	q, err := newQuery(v, OpTypeArrayContainsAny)
	if err != nil {
		qc.err = err
		return qc
	}
	qc.QueryGroup = append(qc.QueryGroup, q)
	return qc
}

// Asc - `Asc`
func (qc *QueryChainer) Asc() *QueryChainer {
	qc.OrderByDirection = firestore.Asc
	return qc
}

// Desc - `Desc`
func (qc *QueryChainer) Desc() *QueryChainer {
	qc.OrderByDirection = firestore.Desc
	return qc
}

// Error - error
func (qc *QueryChainer) Error() error {
	return qc.err
}

// Cursor - query cursor
type Cursor struct {
	isStartAt, isStartAfter bool
	isEndAt, isEndBefore    bool
	values                  []interface{}
}

// StartAt - start at
func (qc *QueryChainer) StartAt(v ...interface{}) *QueryChainer {
	qc.cursor.isStartAt = true
	qc.cursor.values = v
	return qc
}

// StartAt - start after
func (qc *QueryChainer) StartAfter(v ...interface{}) *QueryChainer {
	qc.cursor.isStartAfter = true
	qc.cursor.values = v
	return qc
}

// EndAt - end at
func (qc *QueryChainer) EndAt(v ...interface{}) *QueryChainer {
	qc.cursor.isEndAt = true
	qc.cursor.values = v
	return qc
}

// EndBefore - end before
func (qc *QueryChainer) EndBefore(v ...interface{}) *QueryChainer {
	qc.cursor.isEndBefore = true
	qc.cursor.values = v
	return qc
}

// BuildCursor - build query for cursor
func (qc *QueryChainer) BuildCursorQuery(q firestore.Query) firestore.Query {
	if qc.OrderByDirection == 0 || len(qc.cursor.values) == 0 {
		return q
	}

	if qc.cursor.isStartAt {
		return q.StartAt(qc.cursor.values...)
	}

	if qc.cursor.isStartAfter {
		return q.StartAfter(qc.cursor.values...)
	}

	if qc.cursor.isEndAt {
		return q.EndAt(qc.cursor.values...)
	}

	if qc.cursor.isEndBefore {
		return q.EndBefore(qc.cursor.values...)
	}

	return q
}

// Filter - filters of extra indexer
type Filter struct {
	FilterTypes []FilterType
	Value       interface{}
}

// Filters - using `xim`
func (qc *QueryChainer) Filters(v interface{}, filterTypes ...FilterType) *QueryChainer {
	qc.Filter.Value = v
	if len(filterTypes) == 0 {
		qc.Filter.FilterTypes = append(qc.Filter.FilterTypes, FilterTypeAdd)
		return qc
	}
	qc.Filter.FilterTypes = filterTypes
	return qc
}

// Query - query
type Query struct {
	Operator OpType
	Value    interface{}
}

func newQuery(v interface{}, opType OpType) (*Query, error) {
	dupErr := func(err error) (*Query, error) {
		return nil, xerrors.Errorf("failed to deduplication: %w", err)
	}
	switch x := v.(type) {
	case []bool:
		if err := dedupe.Do(&x); err != nil {
			return dupErr(err)
		}
		v = x
	case []string:
		if err := dedupe.Do(&x); err != nil {
			return dupErr(err)
		}
		v = x
	case []int:
		if err := dedupe.Do(&x); err != nil {
			return dupErr(err)
		}
		v = x
	case []int64:
		if err := dedupe.Do(&x); err != nil {
			return dupErr(err)
		}
		v = x
	case []float64:
		if err := dedupe.Do(&x); err != nil {
			return dupErr(err)
		}
		v = x
	case bool,
		string,
		int,
		int64,
		float64,
		*latlng.LatLng,
		[]*latlng.LatLng,
		*firestore.DocumentRef,
		[]*firestore.DocumentRef,
		map[string]bool,
		map[string]string,
		map[string]int,
		map[string]int64,
		map[string]float64:
		// ok
	case time.Time:
		v = SetLastThreeToZero(x)
	case *time.Time:
		t := SetLastThreeToZero(*x)
		v = &t
	case []time.Time:
		after := make([]time.Time, len(x))
		for n, t := range x {
			after[n] = SetLastThreeToZero(t)
		}
		v = after
	default:
		return nil, xerrors.Errorf("unsupported types: %#v", v)
	}

	q := &Query{
		Operator: opType,
		Value:    v,
	}
	return q, nil
}

// IsSlice - slice judgment
func IsSlice(v interface{}) bool {
	switch v.(type) {
	case []bool, []string, []int, []int64, []float64,
		[]*latlng.LatLng, []*firestore.DocumentRef:
		return true
	}
	return false
}

// IsSlice - slice judgment
func (q *Query) IsSlice() bool {
	return IsSlice(q.Value)
}
