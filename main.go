//  pdf-convert listens on $PORT:/convert for pdfs to convert 
//  docker run -d -e PORT=80 -p 8080:80 <docker image id>
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func checkPath() {
	_, err := exec.LookPath("pdftotext")
	if err != nil {
		log.Fatal("pdftotext not installed")
	}
}

func convertOutput(filename string) (string, error) {
	out, err := exec.Command("pdftotext", filename, "-").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func convertHandler(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal("request", err)
	}
	defer req.Body.Close()
	f, err := ioutil.TempFile("", "pdf_convert")
	if err != nil {
		log.Fatal("Cannot open temp file.", err)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	err = ioutil.WriteFile(f.Name(), buf, 0644)
	if err != nil {
		log.Fatal("Cannot write pdf file.", err)
	}
	out, err := convertOutput(f.Name())
	if err != nil {
		log.Fatal("Conversion failed.", err)
	} else {
		log.Printf("Converted document.")
		fmt.Fprintf(w, out)
	}

}

func main() {
	checkPath()
	http.HandleFunc("/convert", convertHandler)
	port := ":"+os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, nil))
}
