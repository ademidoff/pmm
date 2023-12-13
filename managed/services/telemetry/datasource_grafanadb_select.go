// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package telemetry provides telemetry functionality.
package telemetry

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	// Events, errors and driver for grafana database.
	pmmv1 "github.com/percona-platform/saas/gen/telemetry/events/pmm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type dsGrafanaDBSelect struct {
	l      *logrus.Entry
	config DSConfigGrafanaDB
	db     *sql.DB
}

// check interfaces.
var (
	_ DataSource = (*dsGrafanaDBSelect)(nil)
)

// Enabled flag that determines if the data source is enabled.
func (d *dsGrafanaDBSelect) Enabled() bool {
	return d.config.Enabled
}

// NewDsGrafanaDBSelect makes a new data source to collect grafana metrics.
func NewDsGrafanaDBSelect(config DSConfigGrafanaDB, l *logrus.Entry) DataSource {
	return &dsGrafanaDBSelect{
		l:      l,
		config: config,
	}
}

func (d *dsGrafanaDBSelect) Init(ctx context.Context) error {
	db, err := openGrafanaDBConnection(d.config, d.l)
	if err != nil {
		return err
	}

	d.db = db
	return nil
}

func openGrafanaDBConnection(config DSConfigGrafanaDB, l *logrus.Entry) (*sql.DB, error) {
	var user *url.Userinfo
	if config.UseSeparateCredentials {
		user = url.UserPassword(config.SeparateCredentials.Username, config.SeparateCredentials.Password)
	} else {
		user = url.UserPassword(config.Credentials.Username, config.Credentials.Password)
	}
	uri := url.URL{
		Scheme:   config.DSN.Scheme,
		User:     user,
		Host:     config.DSN.Host,
		Path:     config.DSN.DB,
		RawQuery: config.DSN.Params,
	}
	dsn := uri.String()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a connection pool to PostgreSQL")
	}

	db.SetConnMaxIdleTime(time.Second * 30)
	db.SetConnMaxLifetime(time.Second * 180)
	db.SetMaxIdleConns(1)
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		l.Warnf("Grafana DB is not reachable [%s]: %s", config.DSN.Host, err)
	}

	return db, nil
}

func (d *dsGrafanaDBSelect) FetchMetrics(ctx context.Context, config Config) ([]*pmmv1.ServerMetric_Metric, error) {
	return fetchMetricsFromDB(ctx, d.l, d.config.Timeout, d.db, config)
}

func (d *dsGrafanaDBSelect) Dispose(ctx context.Context) error {
	return d.db.Close()
}