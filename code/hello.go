package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

//START OMIT
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/safe", safe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
	}
	w.Write(b)
}

func safe(w http.ResponseWriter, r *http.Request) {
	Dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	Version, _ := ioutil.ReadFile("/usr/lib/os-release")

	GetUser := exec.Command("whoami")
	
	GetUser.SysProcAttr = &syscall.SysProcAttr{}
  	Username, _ := GetUser.Output()

  	CAP_Check := exec.Command("pscap")
  	CAP_Check.SysProcAttr = &syscall.SysProcAttr{}
  	CAP_Out, _ := CAP_Check.Output()

  	fmt.Fprintf(w, "Dir: %s\nOS Info: \n%s\nUsername: %s\nCapabilities:\n%s\n", string(Dir), string(Version), string(Username), string(CAP_Out))
}

//END OMIT