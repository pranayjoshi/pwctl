package db

import (
	"encoding/json"
	"net/http"
	"pwctl/utils"
)

type DB struct {
	MD_Name                    string  `json:"md_name"`
	MD_Connstr                 string  `json:"md_connstr"`
	MD_IsSuperuser             bool    `json:"md_is_superuser"`
	MD_PresetConfigName        string  `json:"md_preset_config_name"`
	MD_Config                  *string `json:"md_config"`
	MD_IsEnabled               bool    `json:"md_is_enabled"`
	MD_LastModifiedOn          string  `json:"md_last_modified_on"`
	MD_Dbtype                  string  `json:"md_dbtype"`
	MD_IncludePattern          *string `json:"md_include_pattern"`
	MD_ExcludePattern          *string `json:"md_exclude_pattern"`
	MD_CustomTags              *string `json:"md_custom_tags"`
	MD_Group                   string  `json:"md_group"`
	MD_Encryption              string  `json:"md_encryption"`
	MD_HostConfig              *string `json:"md_host_config"`
	MD_OnlyIfMaster            bool    `json:"md_only_if_master"`
	MD_PresetConfigNameStandby *string `json:"md_preset_config_name_standby"`
	MD_ConfigStandby           *string `json:"md_config_standby"`
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) GetDBs(serverAddress string) ([]DB, error) {
	dbURL := "http://" + serverAddress + "/db"

	req, err := http.NewRequest("GET", dbURL, nil)

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

	var dbs []DB
	err = json.NewDecoder(resp.Body).Decode(&dbs)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}
