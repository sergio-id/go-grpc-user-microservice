package postgres

import (
	"database/sql"
	"fmt"
)

// GetPostgresStats returns postgres stats
func GetPostgresStats(dbStats sql.DBStats) string {
	return fmt.Sprintf("MaxOpenConnections: {%v}, OpenConnections: {%v}, InUse: {%v}, Idle: {%v}, WaitCount: {%v}, WaitDuration: {%v}, MaxIdleClosed: {%v}, MaxIdleTimeClosed: {%v}, MaxLifetimeClosed: {%v},",
		dbStats.MaxOpenConnections,
		dbStats.OpenConnections,
		dbStats.InUse,
		dbStats.Idle,
		dbStats.WaitCount,
		dbStats.WaitDuration,
		dbStats.MaxIdleClosed,
		dbStats.MaxIdleTimeClosed,
		dbStats.MaxLifetimeClosed,
	)
}
