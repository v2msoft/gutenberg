package main

import (
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"os"
)

func ExampleNewPDFGenerator() {

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Add one page from an URL
	pdfg.AddPage(wkhtmltopdf.NewPage("https://github.com/SebastiaanKlippert/go-wkhtmltopdf"))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")

}

func main() {

	//Try to load the application configurations. If error is returned stop execution.
	config, err := LoadConfiguration()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(CONFIGURATION_FILE_ERROR_EXIT_STATUS)
	}

	//Set the WkhtmlToPdf path depending on the operating system that is
	//executing the binary file.
	err = SetHtmlToPdfBinaryPath()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(HTML_TO_PDF_CONVERTER_BINARY_ERROR_EXIT_STATUS)
	}

	fmt.Printf("%v", config)

}
