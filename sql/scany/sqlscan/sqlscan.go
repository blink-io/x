package sqlscan

import "github.com/georgysavva/scany/v2/sqlscan"

var (
	DefaultAPI = sqlscan.DefaultAPI

	NotFound = sqlscan.NotFound

	Get         = sqlscan.Get
	ScanAll     = sqlscan.ScanAll
	ScanAllSets = sqlscan.ScanAllSets
	ScanOne     = sqlscan.ScanOne
	ScanRow     = sqlscan.ScanRow
	Select      = sqlscan.Select

	NewAPI        = sqlscan.NewAPI
	NewDBScanAPI  = sqlscan.NewDBScanAPI
	NewRowScanner = sqlscan.NewRowScanner
)

type (
	API = sqlscan.API

	Querier = sqlscan.Querier

	RowScanner = sqlscan.RowScanner
)
