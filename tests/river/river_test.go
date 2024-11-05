package river

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverdatabasesql"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var ctx = context.Background()

type SortArgs struct {
	// Strings is a slice of strings to sort.
	Strings []string `json:"strings"`
}

func (SortArgs) Kind() string { return "sort" }

type SortWorker struct {
	// An embedded WorkerDefaults sets up default methods to fulfill the rest of
	// the Worker interface:
	river.WorkerDefaults[SortArgs]
}

func (w *SortWorker) Work(ctx context.Context, job *river.Job[SortArgs]) error {
	sort.Strings(job.Args.Strings)
	fmt.Printf("Sorted strings: %+v\n", job.Args.Strings)
	return nil
}

func TestRiverWorker_1(t *testing.T) {
	workers := river.NewWorkers()
	// AddWorker panics if the worker is already registered or invalid:
	err := river.AddWorkerSafely(workers, &SortWorker{})
	require.NoError(t, err)

	dbPool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		// handle error
	}

	rc, err := river.NewClient(riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 100},
		},
		Workers: workers,
	})
	require.NoError(t, err)

	// Run the client inline. All executed jobs will inherit from ctx:
	err = rc.Start(ctx)
	require.NoError(t, err)

	defer rc.Stop(ctx)
}

func TestRiver_Bun_1(t *testing.T) {
	workers := river.NewWorkers()
	// AddWorker panics if the worker is already registered or invalid:
	err := river.AddWorkerSafely(workers, &SortWorker{})
	require.NoError(t, err)

	sqlDB, err := sql.Open("pgx", "postgres://localhost/river")
	require.NoError(t, err)

	rc, err := river.NewClient(riverdatabasesql.New(sqlDB), &river.Config{
		Workers: workers,
	})
	require.NoError(t, err)

	bunDB := bun.NewDB(sqlDB, pgdialect.New())

	tx, err := bunDB.BeginTx(ctx, &sql.TxOptions{})
	require.NoError(t, err)

	_, err = rc.InsertTx(ctx, tx.Tx, SortArgs{ // tx.Tx is *sql.Tx
		Strings: []string{
			"whale", "tiger", "bear",
		},
	}, nil)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)

	// Run the client inline. All executed jobs will inherit from ctx:
	err = rc.Start(ctx)
	require.NoError(t, err)

	defer rc.Stop(ctx)
}
