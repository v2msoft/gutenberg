package main

import (
	"runtime"
	"errors"
	"path/filepath"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"os/exec"
)

//SetHtmlToPdfBinaryPath configures the path to the binary which is responsible of converting the HTML files to PDF.
//This path will be different depending on the executing operating system and the CPU's architecture. If the current
//architecture or operating system is not supported an error will be returned.
func SetHtmlToPdfBinaryPath() error{

	//Get the path separator
	separator := string(filepath.Separator)

	//Get the executable base path.
	basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if(err != nil){
		return err
	}else{
		basePath = basePath + separator + "lib"
	}

	//Get the architecture folder name
	archFolderName, err := GetArchitecturePathName()
	if(err != nil){
		return err
	}

	//Get the binary path depending on the executing operating system.
	if(runtime.GOOS == "windows"){
		wkhtmltopdf.SetPath(basePath + separator + "windows" + separator +
			archFolderName + separator + "wkhtmltopdf.exe")
		return nil
	}else if(runtime.GOOS == "linux"){
		wkhtmltopdf.SetPath(basePath + separator + "linux" + separator +
			archFolderName + separator + "bin" + separator + "wkhtmltopdf")
		return nil
	}else if(runtime.GOOS == "darwin"){
		path, err := exec.LookPath("wkhtmltopdf")
		if(err != nil){
			return errors.New("The binary for wkhtmltopdf was not found on your system!")
		}
		wkhtmltopdf.SetPath(path)
		return nil
	}else{
		return errors.New("The operating system " + runtime.GOOS + " is not supported!")
	}

}

//GetArchitecturePathName returns the name of the folder that matches the current platform's architecture. If the
//current architecture is not supported by the application an error object will be returned and the folder path will
//be an empty string.
func GetArchitecturePathName() (string,error){

	//Check if the current platform's architecture is supported.
	if(runtime.GOARCH == "386"){
		return "x86",nil
	}else if(runtime.GOARCH == "amd64"){
		return "x64",nil
	}else{
		return "",errors.New("The architecture " + runtime.GOARCH + " is not supported!")
	}

}