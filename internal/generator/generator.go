package generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func Generate(metadataDir string, outputFile string, includeProperties bool) error {

	var entities []Entity

	fd, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}
	defer fd.Close()

	dir, err := os.Open(metadataDir)
	if err != nil {
		log.Fatalf("Error opening directory: %v", err)
	}
	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, fileName := range fileInfo {
		if !strings.HasSuffix(fileName.Name(), ".json") {
			continue
		}

		file, err := os.Open(dir.Name() + "/" + fileName.Name())
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		var entity Entity
		err = json.NewDecoder(file).Decode(&entity)
		if err != nil {
			log.Fatalf("Error decoding JSON: %v", err)
		}

		entities = append(entities, entity)
	}

	for _, entity := range entities {
		fd.WriteString(fmt.Sprintf("[%s]\n", entity.Name))
		if includeProperties {
			for _, field := range entity.Fields {
				var mandatory string
				if field.Mandatory {
					mandatory = "*"
				} else {
					mandatory = ""
				}
				fd.WriteString(fmt.Sprintf("%s%s\n", mandatory, field.Id))
			}
		}
		fd.WriteString(fmt.Sprintf("\n"))

		for _, relation := range entity.Relations {
			var multiplicity string
			if relation.Type == "array" {
				multiplicity = "*"
			} else {
				multiplicity = "1"
			}
			fd.WriteString(fmt.Sprintf("%s 1--%s %s {label: \"%s\"}\n", entity.Name, multiplicity, relation.Target.Table, relation.Id))
		}
		fd.WriteString(fmt.Sprintf("\n"))
	}

	return nil
}
