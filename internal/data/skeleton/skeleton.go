package skeleton

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	DataSiswaa "go-skeleton-auth/internal/entity/skeleton"
	"go-skeleton-auth/pkg/errors"
	jaegerLog "go-skeleton-auth/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

// Tambahkan query di dalam const
// getAllUser = "GetAllUser"
// qGetAllUser = "SELECT * FROM users"
const (
	getAllData  = "GetAllData"
	qGetAllData = `select * from data_siswa`
)

// Tambahkan query ke dalam key value order agar menjadi prepared statements
//
//	readStmt = []statement{
//		{getAllUser, qGetAllUser},
//	}
var (
	readStmt = []statement{
		{getAllData, qGetAllData},
	}
)

// New ...
func New(db *sqlx.DB, tracer opentracing.Tracer, logger jaegerLog.Factory) Data {
	d := Data{
		db:     db,
		tracer: tracer,
		logger: logger,
	}

	d.initStmt()
	return d
}

func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}

// contoh implementasi ...
// func (d Data) GetShowname(ctx context.Context, movieID string) (string, error) {
// 	var (
// 		showname string
// 		err      error
// 	)

//// WAJIB ADA
// 	if span := opentracing.SpanFromContext(ctx); span != nil {
// 		span := d.tracer.StartSpan("SQL SELECT", opentracing.ChildOf(span.Context()))
// 		span.SetTag("mysql.server", "123.72.156.4")
// 		span.SetTag("mysql.database", "movie")
// 		span.SetTag("mysql.table", "showname")
// 		span.SetTag("mysql.query", "SELECT * FROM movie.showname WHERE movie_id="+movieID)
// 		defer span.Finish()
// 		ctx = opentracing.ContextWithSpan(ctx, span)
// 	}
//// WAJIB ADA

// 	// assumed data fetched from database
// 	showname = "Joni Bizarre Adventure"

//// OPTIONAL, DISARANKAN DIBUAT LOGGINGNYA
// 	d.logger.For(ctx).Info("SQL Query Success", zap.String("showname", showname))

//// WAJIB ADA, INI MERUPAKAN LOGGING KALAU TERJADI ERROR, BISA DIPASANG DI SERVICE DAN HANDLER, TIDAK HANYA DI DATA SAJA
// 	// if err != nil {
// 	// 	d.logger.For(ctx).Error("SQL Query Failed", zap.Error(err))
// 	// 	return showname, err
// 	// }
//// WAJIB ADA

// 	return showname, err
// }

func (d Data) GetDataSiswa(ctx context.Context) ([]DataSiswaa.DataSiswa, error) {
	var (
		datasiswa  DataSiswaa.DataSiswa
		datasiswas []DataSiswaa.DataSiswa
		err        error
	)
	log.Println("data GetDataSiswa object")
	rows, err := d.stmt[getAllData].QueryxContext(ctx)
	if err != nil {
		return datasiswas, errors.Wrap(err, "[DATA] [GetDataSiswa]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&datasiswa); err != nil {
			return datasiswas, errors.Wrap(err, "[DATA] [GetDataSiswa]")
		}
		datasiswas = append(datasiswas, datasiswa)
	}
	return datasiswas, err
}
