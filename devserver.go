package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
	"os"
	"fmt"
	"path/filepath"
	"time"
	"runtime"
	"os/exec"
)

type siteconfig struct  {
	Portnum string 		`json:"portnum"`
	Sitedir string			`json:"site_dir"`
	AppMode bool		`json:"application_mode"`
}

func OpenBrowser(url string){
	//sleep a moment to let the webserver startup first
	time.Sleep( 1 * time.Second )
	//Now launch the proper browser for this OS
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		//Linux, *BSD, etc
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//First load the site config
        var CONFIG siteconfig
	dat, err := ioutil.ReadFile("siteconfig.json")
	if err == nil {
		err = json.Unmarshal(dat, &CONFIG)
		if err != nil { fmt.Printf("Could not parse JSON config!", err) ; os.Exit(1) }
	}
	//Validate Site config
	if(CONFIG.Sitedir == ""){ CONFIG.Sitedir = "dist" }
	if(CONFIG.Portnum == ""){ CONFIG.Portnum = "8844" }

	//Now Start the tiny webserver for this dir on the designated port
	url := "127.0.0.1:"+CONFIG.Portnum
	log.Printf("Starting HTTP Server at %q", url)

	APIHandler := func(w http.ResponseWriter, req *http.Request) {
		path := filepath.Clean(req.URL.EscapedPath())
		if( path == "/"  || path == "" ){ path = "/index.html" }
		if(path == "/quit" && CONFIG.AppMode ){ os.Exit(0) } //got the signal to stop the local app backend
		_, err := os.Stat(CONFIG.Sitedir+path);
		if( err != nil ) {
			path = "/static"+path //look in the static dir for this file instead
			_, err = os.Stat(CONFIG.Sitedir+path);
			if(err != nil){ return } //nothing to do
		}
		filesource, err := os.Open(CONFIG.Sitedir+path)
		if(err != nil){ return }
		defer filesource.Close()
		io.Copy(w, filesource);
	}

	http.HandleFunc("/", APIHandler)
	go OpenBrowser("http://"+url)
	log.Fatal(http.ListenAndServe(url, nil))
}
