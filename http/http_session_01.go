package main

import (
    "log"
    "net/http"
    "html/template"
)

type User struct {
    name    string
    passwd  string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]User{}

func init() {
    tpl = template.Must(template.ParseGlob("tpl/*"))
}

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/test", test)
    http.HandleFunc("/favicon.ico", http.NotFoundHandler())

    err := http.ListenAndServe(":8010", nil)
    if err != nil {
        log.Fatal(""Lsiten Http:", err)
    }
}

func index(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session")
    if err != nil {
        sid := uuid.NewV4()
        c = &http.Cookie{
            Name: "session",
            Value: sid.String(),
        }

        http.SetCookie(w, c)
    }

    var u User
    if  us, ok := dbSession[c.Value]; ok {
        u = dbUser[us]
    }

    if r.Method == http.MethodPost {
        name := r.FormValue("name")
        pwd := r.FormValue("passwd")

        u = User{name, pwd}
        dbSession[c.Value] = us
        dbUsers[us] = u
    }

    tpl.ExecuteTemplate(w, "01_seesion/index.html", u)
}

func test(w http.ResponseWriter, r *http.Request) {
    c, err := r.Cookie("session")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    us, ok := dbSession[c.Value]
    if !ok {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    u := dbUsers[us]
    tpl.ExecuteTemplate(w, "01_seesion/test.html", u)
}
