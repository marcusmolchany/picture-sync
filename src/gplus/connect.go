// Reference: https://developers.google.com/+/quickstart/go
// Reference: https://github.com/googleplus/gplus-quickstart-go

package gplus

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "strings"

    "code.google.com/p/goauth2/oauth"
    "code.google.com/p/google-api-go-client/plus/v1"

    "picture-sync/src/common"
)

const (
    ClientID        = "293255378231-qn21vpjvse30r3ejonc341b2hfle2p07.apps.googleusercontent.com"
    ClientSecret    = "gkpiP6-uEQx_T0l-Imz8eGUu"
    ApplicationName = "gplus-quickstart"
)

// config is the configuration specification supplied to the OAuth package.
var config = &oauth.Config{
    ClientId:     ClientID,
    ClientSecret: ClientSecret,
    // Scope determines which API calls you are authorized to make
    Scope:    "https://www.googleapis.com/auth/plus.login",
    AuthURL:  "https://accounts.google.com/o/oauth2/auth",
    TokenURL: "https://accounts.google.com/o/oauth2/token",
    // Use "postmessage" for the code-flow for server side apps
    RedirectURL: "postmessage",
}

// Token represents an OAuth token response.
type Token struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
    ExpiresIn   int    `json:"expires_in"`
    IdToken     string `json:"id_token"`
}

// ClaimSet represents an IdToken response.
type ClaimSet struct {
    Sub string
}

// connect exchanges the one-time authorization code for a token and common.Stores the
// token in the session
func Connect(w http.ResponseWriter, r *http.Request) *common.AppError {
    // Ensure that the request is not a forgery and that the user sending this
    // connect request is the expected user
    session, err := common.Store.Get(r, "sessionName")
    
    if err != nil {
        log.Println("error fetching session:", err)
        return &common.AppError{err, "Error fetching session", 500}
    }
    
    if r.FormValue("state") != session.Values["state"].(string) {
        m := "Invalid state parameter"
        return &common.AppError{errors.New(m), m, 401}
    }
    
    // Normally, the state is a one-time token; however, in this example, we want
    // the user to be able to connect and disconnect without reloading the page.
    // Thus, for demonstration, we don't implement this best practice.
    // session.Values["state"] = nil

    // Setup for fetching the code from the request payload
    x, err := ioutil.ReadAll(r.Body)
    
    if err != nil {
        return &common.AppError{err, "Error reading code in request body", 500}
    }

    code := string(x)
    accessToken, idToken, err := exchange(code)
    
    if err != nil {
        return &common.AppError{err, "Error exchanging code for access token", 500}
    }
    
    gplusID, err := decodeIdToken(idToken)
    
    if err != nil {
        return &common.AppError{err, "Error decoding ID token", 500}
    }

    // Check if the user is already connected
    storedToken := session.Values["accessToken"]
    storedGPlusID := session.Values["gplusID"]
    
    if storedToken != nil && storedGPlusID == gplusID {
        m := "Current user already connected"
        return &common.AppError{errors.New(m), m, 200}
    }

    // Store the access token in the session for later use
    session.Values["accessToken"] = accessToken
    session.Values["gplusID"] = gplusID
    session.Save(r, w)
    
    return nil
}

// disconnect revokes the current user's token and resets their session
func Disconnect(w http.ResponseWriter, r *http.Request) *common.AppError {
    // Only disconnect a connected user
    session, err := common.Store.Get(r, "sessionName")
    
    if err != nil {
        log.Println("error fetching session:", err)
        
        return &common.AppError{err, "Error fetching session", 500}
    }
    
    token := session.Values["accessToken"]
    
    if token == nil {
        m := "Current user not connected"
        
        return &common.AppError{errors.New(m), m, 401}
    }

    // Execute HTTP GET request to revoke current token
    url := "https://accounts.google.com/o/oauth2/revoke?token=" + token.(string)
    resp, err := http.Get(url)
    
    if err != nil {
        m := "Failed to revoke token for a given user"
        
        return &common.AppError{errors.New(m), m, 400}
    }
    
    defer resp.Body.Close()

    // Reset the user's session
    session.Values["accessToken"] = nil
    session.Save(r, w)
    
    return nil
}

// people fetches the list of people user has shared with this app
func People(w http.ResponseWriter, r *http.Request) *common.AppError {
    session, err := common.Store.Get(r, "sessionName")
    
    if err != nil {
        log.Println("error fetching session:", err)
        
        return &common.AppError{err, "Error fetching session", 500}
    }
    
    token := session.Values["accessToken"]
    
    // Only fetch a list of people for connected users
    if token == nil {
        m := "Current user not connected"
        
        return &common.AppError{errors.New(m), m, 401}
    }

    // Create a new authorized API client
    t := &oauth.Transport{Config: config}
    tok := new(oauth.Token)
    tok.AccessToken = token.(string)
    t.Token = tok
    service, err := plus.New(t.Client())
    
    if err != nil {
        return &common.AppError{err, "Create Plus Client", 500}
    }

    // Get a list of people that this user has shared with this app
    people := service.People.List("me", "visible")
    peopleFeed, err := people.Do()
    
    if err != nil {
        m := "Failed to refresh access token"
        
        if err.Error() == "AccessTokenRefreshError" {
            return &common.AppError{errors.New(m), m, 500}
        }
        
        return &common.AppError{err, m, 500}
    }
    
    w.Header().Set("Content-type", "application/json")
    err = json.NewEncoder(w).Encode(&peopleFeed)
    
    if err != nil {
        return &common.AppError{err, "Convert PeopleFeed to JSON", 500}
    }
    
    return nil
}

// exchange takes an authentication code and exchanges it with the OAuth
// endpoint for a Google API bearer token and a Google+ ID
func exchange(code string) (accessToken string, idToken string, err error) {
    // Exchange the authorization code for a credentials object via a POST request
    addr := "https://accounts.google.com/o/oauth2/token"
    values := url.Values{
        "Content-Type":  {"application/x-www-form-urlencoded"},
        "code":          {code},
        "client_id":     {ClientID},
        "client_secret": {ClientSecret},
        "redirect_uri":  {config.RedirectURL},
        "grant_type":    {"authorization_code"},
    }
    resp, err := http.PostForm(addr, values)
    if err != nil {
        return "", "", fmt.Errorf("Exchanging code: %v", err)
    }
    
    defer resp.Body.Close()

    // Decode the response body into a token object
    var token Token
    err = json.NewDecoder(resp.Body).Decode(&token)
    
    if err != nil {
        return "", "", fmt.Errorf("Decoding access token: %v", err)
    }

    return token.AccessToken, token.IdToken, nil
}

// decodeIdToken takes an ID Token and decodes it to fetch the Google+ ID within
func decodeIdToken (idToken string) (gplusID string, err error) {
    // An ID token is a cryptographically-signed JSON object encoded in base 64.
    // Normally, it is critical that you validate an ID token before you use it,
    // but since you are communicating directly with Google over an
    // intermediary-free HTTPS channel and using your Client Secret to
    // authenticate yourself to Google, you can be confident that the token you
    // receive really comes from Google and is valid. If your server passes the ID
    // token to other components of your app, it is extremely important that the
    // other components validate the token before using it.
    var set ClaimSet
    
    if idToken != "" {
        // Check that the padding is correct for a base64decode
        parts := strings.Split(idToken, ".")
        
        if len(parts) < 2 {
            return "", fmt.Errorf("Malformed ID token")
        }
        
        // Decode the ID token
        b, err := common.Base64Decode(parts[1])
        
        if err != nil {
            return "", fmt.Errorf("Malformed ID token: %v", err)
        }
        
        err = json.Unmarshal(b, &set)
        
        if err != nil {
            return "", fmt.Errorf("Malformed ID token: %v", err)
        }
    }
    return set.Sub, nil
}