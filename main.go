package main

import (
	"bytes"
	"net/http"
)

func getTrafficManagerProfile() {
	json := []byte(`{"grant_type":"client_credentials","client_id":$configuration.client_id,"client_secret":$configuration.client_secret, "resource":"https://management.core.windows.net/"}`)
	authURL := "https://login.microsoftonline.com/$($configuration.subscription)/oauth2/token"
	req, err := http.NewRequest("POST", authURL, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// $uri = "https://management.azure.com/subscriptions/$($Configuration.subscription_id)/resourceGroups/$($Configuration.resource_group)/providers/Microsoft.Network/trafficmanagerprofiles/$($Configuration.traffic_manager_profile)?api-version=2017-05-01"
	//$header = @{"Content-Type" = "application/json"; "Authorization" = "Bearer" + " " + $auth.access_token}
}

func main() {

}
