package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/DixonOrtiz/ApiGateway/api/functions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//GoogleOauthConfig and OauthStateString variables
//Variables that allows to work with OAuth2
var (
	GoogleOauthConfig *oauth2.Config
	OauthStateString  = "pseudo-random"
)

func init() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  functions.GetEnv("GOOGLE_CALLBACK"),
		ClientID:     functions.GetEnv("GOOGLE_CLIENT_ID"),
		ClientSecret: functions.GetEnv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{functions.GetEnv("GOOGLE_SCOPES")},
		Endpoint:     google.Endpoint,
	}
}

//GetUserInfo function
func GetUserInfo(state string, code string) ([]byte, error) {
	if state != OauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
