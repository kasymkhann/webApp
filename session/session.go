package session

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("ger43-sf431-sd"))
