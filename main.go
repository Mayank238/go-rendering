package main

import (
    "html/template"
    "net/http"
    "path"
    "encoding/json"
    "os"
    "io/ioutil"
    "fmt"

)

type Head struct {
    Title  string

}

type form struct{

 }

 // type display struct {

 // }

 type Card struct {
    Name       string `json:"name"`
    Email      string `json:"email"`
    Phone_no   string `json:Phone_no`
    Password   string `json:password`
}

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/form", Form)
    http.HandleFunc("/display", Display)
    http.HandleFunc("/list", list)
    http.ListenAndServe(":" + os.Getenv("PORT"), nil)

}

func Home(w http.ResponseWriter, r *http.Request) {

    title := Head{"Welcome"}

    fp := path.Join("templet", "index.html")
    tmpl, err := template.ParseFiles(fp)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, title); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }


}

func Form(w http.ResponseWriter, r *http.Request) {
    var f form
    t,_ := template.ParseFiles("form.html")
    t.Execute(w, f)

}

func Display(w http.ResponseWriter, r *http.Request) {
    //var d display
    data,err := ioutil.ReadFile("deck.json")
    //fmt.Println(data)

    var Dataobj []Card
    err = json.Unmarshal(data,&Dataobj)
    if err != nil {
       fmt.Println(err)
    }

    //fmt.Println(Dataobj)

    d:= Card{
        Name: r.FormValue("name"),
        Email: r.FormValue("email"),
        Phone_no:r.FormValue("p_no"),
        Password: r.FormValue("pwd"),

    }


    Dataobj = append(Dataobj, d)
    //fmt.Println(d)

    // card := new(Card)
    // card.Name = r.FormValue("name")
    // card.Email = r.FormValue("email")
    // card.Phone_no =r.FormValue("p_no")
    // card.Password = r.FormValue("pwd")

    d2,_ := json.MarshalIndent(Dataobj,""," ")
    err = ioutil.WriteFile("deck.json", []byte(d2), 0644)
    if err != nil{
        fmt.Println(err)
    }

    t,_ := template.ParseFiles("displaydata.html")
    t.Execute(w, d)

    //  f, err := os.OpenFile("deck.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    // if err != nil {
    //     http.Error(w, err.Error(), 500)
    //     return
    // }




    // b, err := json.Marshal(card)
    // if err != nil {
    //     http.Error(w, err.Error(), 500)
    //     return
    // }

    // f.Write(b)
    // f.Close()

}

func list(w http.ResponseWriter,r *http.Request) {



      f, err := os.OpenFile("deck.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
        if err != nil {
        fmt.Println("Error")
       }
       defer f.Close()

    bytevalue,_ := ioutil.ReadAll(f)
    //fmt.Println(bytevalue)
    var Res []Card
     json.Unmarshal(bytevalue, &Res)
       fmt.Println(Res)
     t, _ := template.ParseFiles("list.html")
     t.Execute(w, Res)

}
