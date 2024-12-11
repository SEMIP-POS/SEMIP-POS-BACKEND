package healthCheck

import "context"

func (r *healthCheck) CheckDB(ctx context.Context) error {
	return r.db.Ping()
}
