package auth

import (
	"fmt"
	"io"
	"net/http"
	"pwctl/utils"
	"strings"

	"github.com/urfave/cli/v2"
)

func Login(c *cli.Context) error {

	serverAddress := c.String("c")
	username := c.String("user")
	password := c.String("password")

	loginURL := "http://" + serverAddress + "/login"
	payload := strings.NewReader(`{
        "user": "` + username + `",
        "password": "` + password + `"
    }`)

	req, _ := http.NewRequest("POST", loginURL, payload)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode == 200 {
		utils.StoreToken(string(body))
		utils.StoreTokenFile(string(body))
		fmt.Println("Login successful")
	} else {
		return fmt.Errorf("status code: %d\n", res.StatusCode)
	}

	return nil
}
