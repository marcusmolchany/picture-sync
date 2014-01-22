package common

import (
    "encoding/base64"
    "crypto/rand"
    "github.com/gorilla/sessions"
)

type AppError struct {
    Err     error
    Message string
    Code    int
}

var Store = sessions.NewCookieStore([]byte(RandomString(32)))

// randomString returns a random string with the specified length
func RandomString(length int) (str string) {
    b := make([]byte, length)
    rand.Read(b)

    return base64.StdEncoding.EncodeToString(b)
}

func Base64Decode(s string) ([]byte, error) {
    // add back missing padding
    switch len(s) % 4 {
        case 2:
            s += "=="
        case 3:
            s += "="
    }

    return base64.URLEncoding.DecodeString(s)
}