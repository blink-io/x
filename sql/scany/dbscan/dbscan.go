package dbscan

import "github.com/georgysavva/scany/v2/dbscan"

var (
	DefaultAPI = dbscan.DefaultAPI

	ErrNotFound = dbscan.ErrNotFound

	NotFound = dbscan.NotFound

	ScanAll         = dbscan.ScanAll
	ScanAllSets     = dbscan.ScanAllSets
	ScanOne         = dbscan.ScanOne
	ScanRow         = dbscan.ScanRow
	SnakeCaseMapper = dbscan.SnakeCaseMapper

	NewAPI                  = dbscan.NewAPI
	WithAllowUnknownColumns = dbscan.WithAllowUnknownColumns
	WithColumnSeparator     = dbscan.WithColumnSeparator
	WithFieldNameMapper     = dbscan.WithFieldNameMapper
	WithScannableTypes      = dbscan.WithScannableTypes

	NewRowScanner = dbscan.NewRowScanner
)

type (
	API = dbscan.API

	NameMapperFunc = dbscan.NameMapperFunc

	RowScanner = dbscan.RowScanner

	Rows = dbscan.Rows
)
