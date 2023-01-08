package main

import (
	"fmt"
	"log"
	"os"
)

func createXMLFile() {
	err := os.Mkdir("Archive", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Create a file
	file, err := os.Create("Archive/asset_data.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("<Data>\n"))
	_, err = file.WriteString(fmt.Sprintf("\t<Time>\n\t\t%s\n\t</Time>\n", getTime().Result.Rfc1123))
	_, err = file.WriteString(fmt.Sprintf("\t<SystemStatus>\n\t\t%s\n\t</SystemStatus>\n", getStatus().Result.Status))
	_, err = file.WriteString(fmt.Sprintf("</Data>"))

	if err != nil {
		log.Fatal(err)
	}
}
