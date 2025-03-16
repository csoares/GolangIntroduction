package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Config represents application configuration
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Logging struct {
		Level  string `json:"level"`
		Output string `json:"output"`
	} `json:"logging"`
}

// Person represents a record in the CSV file
type Person struct {
	Name string
	Age  string
	City string
}

func main() {
	// Create a directory for our example files
	exampleDir := "example_files"
	os.Mkdir(exampleDir, 0755)

	// 1. Basic File I/O Example
	fmt.Println("=== Basic File I/O Example ===")
	basicFileExample(exampleDir)

	// 2. CSV File Example
	fmt.Println("\n=== CSV File Example ===")
	csvFileExample(exampleDir)

	// 3. JSON Config Example
	fmt.Println("\n=== JSON Config Example ===")
	configFileExample(exampleDir)
}

// Basic File I/O Example Functions
func basicFileExample(exampleDir string) {
	filePath := filepath.Join(exampleDir, "sample.txt")
	content := "Hello, Go File I/O!\nThis is a sample text file.\n"

	// Writing to a file
	err := writeFile(filePath, content)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	// Reading from a file
	readContent, err := readFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("File contents:")
	fmt.Println(readContent)

	// Appending to a file
	appendContent := "This line was appended!\n"
	err = appendToFile(filePath, appendContent)
	if err != nil {
		fmt.Printf("Error appending to file: %v\n", err)
		return
	}

	// Reading the file again
	readContent, err = readFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("\nUpdated file contents:")
	fmt.Println(readContent)
}

func writeFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func appendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// CSV File Example Functions
func csvFileExample(exampleDir string) {
	filePath := filepath.Join(exampleDir, "people.csv")

	// Sample data
	people := []Person{
		{"John Doe", "30", "New York"},
		{"Jane Smith", "25", "Los Angeles"},
		{"Bob Johnson", "35", "Chicago"},
	}

	// Write CSV file
	err := writeCSV(filePath, people)
	if err != nil {
		fmt.Printf("Error writing CSV: %v\n", err)
		return
	}

	// Read and display CSV file
	readPeople, err := readCSV(filePath)
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return
	}

	fmt.Println("CSV Contents:")
	for _, person := range readPeople {
		fmt.Printf("Name: %s, Age: %s, City: %s\n", person.Name, person.Age, person.City)
	}
}

func writeCSV(filePath string, people []Person) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "Age", "City"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data
	for _, person := range people {
		row := []string{person.Name, person.Age, person.City}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func readCSV(filePath string) ([]Person, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read data
	var people []Person
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}
		people = append(people, Person{
			Name: record[0],
			Age:  record[1],
			City: record[2],
		})
	}

	return people, nil
}

// JSON Config Example Functions
func configFileExample(exampleDir string) {
	filePath := filepath.Join(exampleDir, "config.json")

	// Create sample configuration
	config := Config{}
	config.Server.Host = "localhost"
	config.Server.Port = 8080
	config.Database.Host = "db.example.com"
	config.Database.Port = 5432
	config.Database.Username = "admin"
	config.Database.Password = "secret"
	config.Logging.Level = "info"
	config.Logging.Output = "app.log"

	// Write configuration to file
	err := writeConfig(filePath, config)
	if err != nil {
		fmt.Printf("Error writing config: %v\n", err)
		return
	}

	// Read configuration from file
	readConfig, err := readConfig(filePath)
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return
	}

	// Display configuration
	fmt.Println("Configuration:")
	fmt.Printf("Server: %s:%d\n", readConfig.Server.Host, readConfig.Server.Port)
	fmt.Printf("Database: %s:%d\n", readConfig.Database.Host, readConfig.Database.Port)
	fmt.Printf("Logging: %s -> %s\n", readConfig.Logging.Level, readConfig.Logging.Output)
}

func writeConfig(filePath string, config Config) error {
	data, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func readConfig(filePath string) (Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
