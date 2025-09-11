package envuse

import (
	"os"
	"fmt"
	"bufio"
)

type OpenFileError struct {
	File string
	Err error
}
func (e OpenFileError) Error() string {
	return fmt.Sprintf("failed to open env file %s: %v", e.File, e.Err)
}

type ReadFileError struct {
	File string
	Err error
}
func (e ReadFileError) Error() string {
	return fmt.Sprintf("failed to read line from env file %s: %v", e.File, e.Err)
}

type EnvKeyError struct {
	Key string
}
func (e EnvKeyError) Error() string {
	return fmt.Sprintf("env key %s does not exist", e.Key)
}

var envMap map[string]string = map[string]string{}

func LoadEnvFile(fileName string) error {
	file, errFile := os.Open(fileName)
	if errFile != nil { return OpenFileError{ File: fileName, Err: errFile } }
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	var fileContent []string
	for fileScanner.Scan() {
		fileContent = append(fileContent, fileScanner.Text()) 
	}
	fileScannerErr := fileScanner.Err()
	if fileScannerErr != nil { return ReadFileError{ File: fileName, Err: fileScannerErr } }
	fmt.Printf("Env file content: %v", fileContent)
	return nil
}

func GetEnv(envKey string) (string, error) {
	value, ok := envMap[envKey]
	if !ok { return "", EnvKeyError{ Key: envKey } }
	return value, nil
}
