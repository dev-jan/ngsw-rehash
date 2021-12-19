package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	ConfigVersion             int                    `json:"configVersion"`
	Timestamp                 int                    `json:"timestamp"`
	Index                     string                 `json:"index"`
	AssetGroups               interface{}            `json:"assetGroups"`
	DataGroups                interface{}            `json:"dataGroups"`
	HashTable                 map[string]interface{} `json:"hashTable"`
	NavigationUrls            interface{}            `json:"navigationUrls"`
	NavigationRequestStrategy string                 `json:"navigationRequestStrategy"`
}

func recreateHashes(ngswConfigPath string) {
	jsonFile, err := os.Open(ngswConfigPath)
	if err != nil {
		log.Fatal("Error while opening ngsw.json file! error: ", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	baseDir := filepath.Dir(ngswConfigPath)

	var result Config
	json.Unmarshal([]byte(byteValue), &result)

	for filename, hash := range result.HashTable {
		file, err := os.Open(baseDir + filename)
		if err != nil {
			log.Fatal("Cannot find file", filename, " from hashtable! ", err)
		}
		defer file.Close()

		sha1Hash := sha1.New()
		if _, err := io.Copy(sha1Hash, file); err != nil {
			log.Fatal("Error while reading file ", filename, ": ", err)
		}

		fileHashSum := fmt.Sprintf("%x", sha1Hash.Sum(nil))
		if fileHashSum != hash {
			fmt.Println(filename, " has a new hash, going to update it:")
			fmt.Println(" ", hash, " => ", fileHashSum)
			result.HashTable[filename] = fileHashSum
		}
	}
	file, _ := json.MarshalIndent(result, "", "  ")
	err = ioutil.WriteFile(ngswConfigPath, file, 0644)
	if err != nil {
		log.Fatal("Error while writing ngsw.json file! error: ", err)
	}
}

func main() {
	fmt.Println("== ngsw-rehash ==")

	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], " [path to ngsw.json]")
		log.Fatal("Path to ngsw.json must be provided as argument")
	}

	ngswConfigPath := os.Args[1]
	recreateHashes(ngswConfigPath)
}
