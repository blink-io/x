package pgxscan

import "github.com/georgysavva/scany/v2/pgxscan"

var (
	DefaultAPI = pgxscan.DefaultAPI

	NotFound = pgxscan.NotFound

	Get     = pgxscan.Get
	ScanAll = pgxscan.ScanAll
	ScanOne = pgxscan.ScanOne
	ScanRow = pgxscan.ScanRow
	Select  = pgxscan.Select

	NewAPI         = pgxscan.NewAPI
	NewDBScanAPI   = pgxscan.NewDBScanAPI
	NewRowScanner  = pgxscan.NewRowScanner
	NewRowsAdapter = pgxscan.NewRowsAdapter
)

type (
	API = pgxscan.API

	Querier = pgxscan.Querier

	RowScanner = pgxscan.RowScanner

	RowsAdapter = pgxscan.RowsAdapter
)
