package main

import(
	"fmt"
	// "html"
    "log"
    "net/http"
    "os"
)

func helloworld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
    // w.Header().Set("Content-Type", "text/html; charset=ascii")
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1
}

func main (){
	http.HandleFunc("/", helloworld)
    port := os.Getenv("PORT")

  if len(port) == 0 {
	port = "5000"
  }

  if err := http.ListenAndServe(":"+port, nil); err != nil {
      log.Fatal(err)
  }
}