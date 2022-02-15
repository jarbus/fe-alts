package main

import (
    "strings"
    "fmt"
    "log"
    "net/http"
    "os"
)


func removeWWW(url string) string { return strings.Replace(url, "www.", "", 1) }
func addHTTP(url string) string {
    if !strings.Contains(url, "http"){
        return fmt.Sprintf("https://%s", url)
    }
    return url
}

func twitter2nitter(url string, twitter_pattern string) string{
    return strings.Replace(url, twitter_pattern, os.Getenv("NITTER_URL"), 1)
}
func medium2scribe(url string, pattern string) string {
    return fmt.Sprintf("https://%s/%s", os.Getenv("SCRIBE_URL"), url)
}

func reddit2teddit(url string, reddit_pattern string) string {
    return strings.Replace(url, reddit_pattern, os.Getenv("TEDDIT_URL"), 1)
}

func youtube2invidious(url string, youtube_pattern string) string {
    return strings.Replace(url, youtube_pattern, os.Getenv("INVIDIOUS_URL"), 1)
}

func insta2biblio(url string, insta_pattern string) string {
    if strings.Contains(url, "/p/") {
        return strings.Replace(url, insta_pattern, os.Getenv("BIBLIOGRAM_URL"), 1)
    }
    return strings.Replace(url, insta_pattern, fmt.Sprintf("%s/u", os.Getenv("BIBLIOGRAM_URL")), 1)
}

var url_pairs = map[string]func(string, string)string{
    "medium.com":            medium2scribe,
    "mobile.twitter.com":    twitter2nitter,
    "twitter.com":           twitter2nitter,
    "reddit.com":            reddit2teddit,
    "t.co":                  twitter2nitter,
    "m.youtube.com":         youtube2invidious,
    "youtube.com":           youtube2invidious,
    "instagram.com":         insta2biblio,
}


func redirectURL(w http.ResponseWriter, r *http.Request, url string){
    for pattern, handler := range url_pairs{
        if (strings.Contains(url, pattern)) {
            new_url  :=  handler(url, pattern)
            new_url  = removeWWW(new_url)
            new_url  = addHTTP(new_url)
            http.Redirect(w, r, new_url, http.StatusSeeOther)
            return
        }
    }
    http.ServeFile(w, r, "index.html")
}


func main() {

    fmt.Printf("Server Started.")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        if err := r.ParseForm(); err != nil { log.Fatal(err) }

        // Path is either the string after / in the url
        // or whatever is entered in the search box
        path := r.URL.Path[1:]
        if r.URL.RawQuery != "" {
            path = path + "?" + r.URL.RawQuery
        }
        if r.Form.Get("website") != "" {
            path = r.Form.Get("website")
        }
        redirectURL(w, r, path)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}
