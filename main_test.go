package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func getHashOfFile(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func TestRehash(t *testing.T) {
	// arrange
	const testfile = "./testdata/unittestfile-ngsw.json"
	if _, err := os.Stat(testfile); err == nil {
		os.Remove(testfile)
	}

	bytes, err := ioutil.ReadFile("./testdata/ngsw.json")
	if err != nil {
		t.Errorf("Cannot read test template ngsw.json file! error: %s", err)
	}
	err = ioutil.WriteFile(testfile, bytes, 0644)
	if err != nil {
		t.Errorf("Cannot write test ngsw.json to %s error: %s", testfile, err)
	}

	// act
	recreateHashes(testfile)

	// check if modified ngsw.json is correct
	actualHash := getHashOfFile(testfile)
	targetHash := getHashOfFile("./testdata/unittest_expected_output.json")

	if actualHash != targetHash {
		t.Errorf("Actualfile is not equal to target file, please check!")
	}

	// remove testfile after successful test
	os.Remove(testfile)
}
