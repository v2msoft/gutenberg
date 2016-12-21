package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func createExampleFile(path string) (*Configuration, error) {

	//Get the example JSON string from Configuration struct.
	configJson, err := json.Marshal(Configuration{})
	if err != nil {
		return nil, err
	}

	//Create the configuration file with example JSON.
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	//Close the file descriptor when leaving the method.
	defer file.Close()

	//Write the example JSON into configuration file.
	writer := bufio.NewWriter(file)
	writer.WriteString(string(configJson))
	writer.Flush()

	//Return an error notifying that parameters on configuration file may be changed
	return nil, errors.New("A new configuration file has been created. Please, review it to validate the" +
		" the configuration options and start the service again.")
}

func readConfigurationFile(path string) (*Configuration, error) {

	//Open the configuration file.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Configuration

	//Try to un-marshall the configuration file.
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	//Return a pointer to the configuration struct that has been filled in with the
	//configuration file information.
	return &config, nil

}

//LoadConfiguration reads or creates a configuration file for the Gutenberg service. If the file is created for the
//first time, an error will be returned to notify that default values have been written into the configuration file.
func LoadConfiguration() (*Configuration, error) {

	//Get the operating system separator
	separator := string(filepath.Separator)

	//Get the application base path.
	basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}

	//Check if the configuration file doesn't exist. If it doesn't exist, create a new one based on the Configuration
	//struct template.
	if _, err := os.Stat(basePath + separator + CONFIGURATION_FILE_NAME); os.IsNotExist(err) {
		return createExampleFile(basePath + separator + CONFIGURATION_FILE_NAME)
	} else {
		return readConfigurationFile(basePath + separator + CONFIGURATION_FILE_NAME)
	}

}
