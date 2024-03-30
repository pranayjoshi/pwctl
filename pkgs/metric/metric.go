package metric

import (
	"encoding/json"
	"net/http"
	"pwctl/utils"
	"time"
)

type Metric struct {
	M_MID            int       `db:"m_id"`
	M_MName          string    `db:"m_name"`
	M_PGVersionFrom  float64   `db:"m_pg_version_from"`
	M_SQL            string    `db:"m_sql"`
	M_Comment        string    `db:"m_comment"`
	M_IsActive       bool      `db:"m_is_active"`
	M_IsHelper       bool      `db:"m_is_helper"`
	M_LastModifiedOn time.Time `db:"m_last_modified_on"`
	M_MasterOnly     bool      `db:"m_master_only"`
	M_StandbyOnly    bool      `db:"m_standby_only"`
	M_ColumnAttrs    string    `db:"m_column_attrs"` // jsonb type can be represented as string in Go
	M_SQLSU          string    `db:"m_sql_su"`
}

func NewMetric() *Metric {
	return &Metric{}
}

func (db *Metric) GetDBs(serverAddress string) ([]Metric, error) {
	metricURL := "http://" + serverAddress + "/metrics"

	req, err := http.NewRequest("GET", metricURL, nil)

	if err != nil {
		return nil, err
	}
	token, _ := utils.GetTokenFile()
	req.Header.Set("Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var metrics []Metric
	err = json.NewDecoder(resp.Body).Decode(&metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
