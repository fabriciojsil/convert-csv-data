package service

import (
	"os"
	"testing"
)

func TestParserService(t *testing.T) {
	tearDown := setupTest(t)
	defer tearDown(t)

	t.Run("Parsing a csv doesn't exists", func(t *testing.T) {
		format := "json"
		filePath := "../data/data.csv"
		savedPath, err := DeliveryFile(format, filePath, "", "")

		if err == nil {
			t.Errorf("Must return err %v", err)
		}

		if savedPath != "" {
			t.Errorf("Expect saved patch %v  Actual %v", "", savedPath)
		}
	})

	t.Run("Parsing to Json", func(t *testing.T) {
		expectedFilePathSaved := "../../data/data.json"
		format := "json"
		filePath := "../../data/data.csv"
		savedPath, err := DeliveryFile(format, filePath, "", "")

		if err != nil {
			t.Errorf("Mustn't return err %v", err)
		}

		if savedPath != expectedFilePathSaved {
			t.Errorf("Expect saved patch %v  Actual %v", expectedFilePathSaved, savedPath)
		}
	})

	t.Run("Parsing to xml", func(t *testing.T) {
		expectedFilePathSaved := "../../data/data.xml"
		format := "xml"
		filePath := "../../data/data.csv"
		savedPath, err := DeliveryFile(format, filePath, "", "")

		if err != nil {
			t.Errorf("Mustn't return err %v", err)
		}

		if savedPath != expectedFilePathSaved {
			t.Errorf("Expect saved patch %v  Actual %v", expectedFilePathSaved, savedPath)
		}
	})

	t.Run("Parsing to an format not suported", func(t *testing.T) {
		expectedFilePathSaved := "../../data/data.json"
		format := "html"
		filePath := "../../data/data.csv"
		savedPath, err := DeliveryFile(format, filePath, "", "")

		if err != nil {
			t.Errorf("Mustn't return err %v", err)
		}

		if savedPath != expectedFilePathSaved {
			t.Errorf("Expect saved patch %v  Actual %v", expectedFilePathSaved, savedPath)
		}
	})
}

func setupTest(t *testing.T) func(t *testing.T) {
	t.Log("ensure there is no parsed files")
	os.Remove("../../data/data.json")
	os.Remove("../../data/data.xml")

	return func(t *testing.T) {
		os.Remove("../../data/data.json")
		os.Remove("../../data/data.xml")
		t.Log("removing parsed files")
	}
}
