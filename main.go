package main

import (
    "html/template"
    "log"
    "net/http"
    "net/url"
    "picture-sync/src/gplus"
    "picture-sync/src/common"
)

func main() {
    fileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/img/", http.FileServer(http.Dir("img")))
    http.Handle("/img/", fileServer)

    // Register a handler for our API calls
    http.Handle("/connect", appHandler(gplus.Connect))
    http.Handle("/disconnect", appHandler(gplus.Disconnect))
    http.Handle("/people", appHandler(gplus.People))

    // Serve the index.html page
    http.Handle("/", appHandler(indexPage))
    err := http.ListenAndServe(":4567", nil)
    
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// index sets up a session for the current user and serves the index page
func indexPage(w http.ResponseWriter, r *http.Request) *common.AppError {
    // This check prevents the "/" handler from handling all requests by default
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return nil
    }

    // Create a state token to prevent request forgery and common.Store it in the session
    // for later validation
    session, err := common.Store.Get(r, "sessionName")
    
    if err != nil {
        log.Println("error fetching session:", err)
        // Ignore the initial session fetch error, as Get() always returns a
        // session, even if empty.
        //return &common.AppError{err, "Error fetching session", 500}
    }
    
    state := common.RandomString(64)
    session.Values["state"] = state
    session.Save(r, w)
    
    stateURL := url.QueryEscape(session.Values["state"].(string))

    // Fill in the missing fields in index.html
    var data = struct {
        ApplicationName, ClientID, State string
    }{gplus.ApplicationName, gplus.ClientID, stateURL}

    // Render and serve the HTML
    // indexTemplate is the HTML template we use to present the index page.
    var indexTemplate = template.Must(template.ParseFiles("index.html"))
    
    err = indexTemplate.Execute(w, data)
    
    if err != nil {
        log.Println("error rendering template:", err)
        return &common.AppError{err, "Error rendering template", 500}
    }
    
    return nil
}

// appHandler is to be used in error handling
type appHandler func(http.ResponseWriter, *http.Request) *common.AppError

// serveHTTP formats and passes up an error
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if e := fn(w, r); e != nil { // e is *common.AppError, not os.Error.
        log.Println(e.Err)
        http.Error(w, e.Message, e.Code)
    }
}