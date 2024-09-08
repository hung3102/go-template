// Package httpstatuscode - HTTPStatusCodeを持ってるか判別するためのパッケージ
package httpstatuscode

// Interface - HTTPStatusCodeを持ってる型
type Interface interface {
	StatusCode() int
}
