package gplus

import (
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "net/http"
    "code.google.com/p/goauth2/oauth"
)

var (
    clientId     = flag.String("id", "320731813162.apps.googleusercontent.com", "Client ID")
    clientSecret = flag.String("secret", "JDKYH4I4AOpHkT18zdUlvSay", "Client Secret")
    scope        = flag.String("scope", "https://www.googleapis.com/auth/userinfo.profile", "OAuth scope")
    redirectURL  = flag.String("redirect_url", "http://localhost/", "Redirect URL")
    authURL      = flag.String("auth_url", "https://accounts.google.com/o/oauth2/auth", "Authentication URL")
    tokenURL     = flag.String("token_url", "https://accounts.google.com/o/oauth2/token", "Token URL")
    requestURL   = flag.String("request_url", "https://www.googleapis.com/oauth2/v1/userinfo", "API request")
    code         = flag.String("code", "", "Authorization Code")
    cachefile    = flag.String("cache", "cache.json", "Token cache file")
)

const usageMsg = `
To obtain a request token you must specify both -id and -secret.

To obtain Client ID and Secret, see the "OAuth 2 Credentials" section under
the "API Access" tab on this page: https://code.google.com/apis/console/

Once you have completed the OAuth flow, the credentials should be stored inside
the file specified by -cache and you may run without the -id and -secret flags.
`

func Connect() (string) {
    flag.Parse()

        // Set up a configuration.
    config := &oauth.Config{
        ClientId:     *clientId,
        ClientSecret: *clientSecret,
        RedirectURL:  *redirectURL,
        Scope:        *scope,
        AuthURL:      *authURL,
        TokenURL:     *tokenURL,
        TokenCache:   oauth.CacheFile(*cachefile),
    }

    // Set up a Transport using the config.
    transport := &oauth.Transport{Config: config}

    // Try to pull the token from the cache; if this fails, we need to get one.
    token, err := config.TokenCache.Token()

    if err != nil {
        if *clientId == "" || *clientSecret == "" {
            flag.Usage()
            fmt.Fprint(os.Stderr, usageMsg)
            os.Exit(2)
        }

        if *code == "" {
            // Get an authorization code from the data provider.
            // ("Please ask the user if I can access this resource.")
            url := config.AuthCodeURL("")
            fmt.Println("Visit this URL to get a code, then run again with -code=YOUR_CODE\n")
            fmt.Println(url)
            return url
        }
        
        // Exchange the authorization code for an access token.
        // ("Here's the code you gave the user, now give me a token!")
        token, err = transport.Exchange(*code)
        if err != nil {
            log.Fatal("Exchange:", err)
        }
        
        // (The Exchange method will automatically cache the token.)
        fmt.Printf("Token is cached in %v\n", config.TokenCache)
    }

    // Make the actual request using the cached token to authenticate.
    // ("Here's the token, let me in!")
    transport.Token = token

    // Make the request.
    r, err := transport.Client().Get(*requestURL)
    
    if err != nil {
        log.Fatal("Get:", err)
    }
    
    defer r.Body.Close()

    // Write the response to standard output.
    io.Copy(os.Stdout, r.Body)

    // Send final carriage return, just to be neat.
    fmt.Println()

    return "fail"
}

func GetRequest(url string) {
    resp, err := http.Get(url)

    fmt.Printf("Response: %v\n", resp)
    fmt.Printf("Error: %v\n", err)
}