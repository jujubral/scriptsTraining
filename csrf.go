package main

import (
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // criar uma nova solicitação HTTP DELETE
        req, err := http.NewRequest(http.MethodDelete, "http://localhost:5127/api/User/13e1c5f8-30b3-4231-a4f4-c787c68148f2", strings.NewReader(""))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // redirecionar para o YouTube
        http.Redirect(w, r, "https://www.youtube.com/watch?v=H0AMHM3EUeA", http.StatusSeeOther)
    })

    http.ListenAndServe(":8080", nil)
}
