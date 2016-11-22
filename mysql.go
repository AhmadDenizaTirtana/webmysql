package main

import "database/sql"
import "fmt"
import "html/template"
import "log"
import "net/http"
import _ "github.com/go-sql-driver/mysql"
import "os"

func connect() *sql.DB {
	var db, err = sql.Open("mysql", "root:@/pengweb")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping database unsuccessful, we may not use db connection")
	} else {
		fmt.Println("Ping database successful, we may use db connection")
	}

	return db
}

func showPerson(w http.ResponseWriter, r *http.Request) {

	var db = connect()
	defer db.Close()

	var id, nama string

	t, _ := template.New("t").Parse(indeks)

	rows, err := db.Query("select * from mahasiswa;")
	if err != nil {
		log.Fatal("Err: ", err)
	}
	for rows.Next() {
		rows.Scan(&id, &nama)
		data := map[string]string{
			"id": id, "nama": nama,
		}
		t.Execute(w, data)
	}
}

const indeks = `<table width='100%' border='1'><tr><td width='30%'>{{.id}}</td><td width='70%'>{{.nama}}</td></tr></table>`

func main() {

	http.HandleFunc("/", showPerson) // set router
	fmt.Println("running server via localhost.....")
	http.ListenAndServe(":9090", nil) // set listen port

}
