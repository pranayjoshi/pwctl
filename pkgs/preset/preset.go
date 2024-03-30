package preset

import (
	"encoding/json"
	"net/http"
	"pwctl/utils"
)

type Preset struct {
	PC_Name           string  `json:"pc_name"`
	PC_IsSuperuser    bool    `json:"pc_description"`
	PC_Config         *string `json:"pc_config"`
	PC_LastModifiedOn string  `json:"pc_last_modified_on"`
}

func NewPreset() *Preset {
	return &Preset{}
}

func (preset *Preset) GetPresets(serverAddress string) ([]Preset, error) {
	presetURL := "http://" + serverAddress + "/preset"

	req, err := http.NewRequest("GET", presetURL, nil)

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

	var presets []Preset
	err = json.NewDecoder(resp.Body).Decode(&presets)
	if err != nil {
		return nil, err
	}

	return presets, nil
}
