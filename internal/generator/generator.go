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

	var relationships = make(map[string][]Relation)

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

		entity := ParseMetadata(dir.Name() + "/" + fileName.Name())

		entities = append(entities, entity)
	}

	for _, entity := range entities {
		fd.WriteString(fmt.Sprintf("[%s]\n", entity.Name))
		if includeProperties {
			for _, field := range entity.Fields {
				s, err := generateSingleProperty(field, entity)
				if err != nil {
					log.Fatalf("Error generating property: %v", err)
				}
				fd.WriteString(s)
			}
		}
		fd.WriteString(fmt.Sprintf("\n"))

		relationships[entity.Name] = entity.Relations

		for _, relation := range entity.Relations {
			s, err := generateSingleRelation(relation, entity)
			if err != nil {
				log.Fatalf("Error generating relation: %v", err)
			}
			fd.WriteString(s)
		}
		fd.WriteString(fmt.Sprintf("\n"))
	}

	// for sKey, sourceRels := range relationships {
	// 	for _, sourceRel := range sourceRels {
	// 		for i, targetRel := range relationships[sourceRel.Target.Table] {
	// 			if sKey == sourceRel.Target.Table {
	// 				// Skip self-references
	// 				log.Printf("Skipping %s", sourceRel.Id)
	// 				var multiplicity string
	// 				if sourceRel.Type == "array" {
	// 					multiplicity = "*"
	// 				} else {
	// 					multiplicity = "1"
	// 				}
	// 				fd.WriteString(fmt.Sprintf("%s 1--%s %s {label: \"%s\"}\n", sKey, multiplicity, sourceRel.Target.Table, sourceRel.Id))
	// 				continue
	// 			}
	// 			if sKey == targetRel.Target.Table {
	// 				log.Printf("Found reverse relationship %s (%s) between %s and %s", sourceRel.Id, targetRel.Id, sKey, sourceRel.Target.Table)
	// 				relationships[sourceRel.Target.Table] = append(relationships[sourceRel.Target.Table][:i], relationships[sourceRel.Target.Table][i+1:]...)

	// 				var multiplicity string
	// 				if sourceRel.Type == "array" {
	// 					multiplicity = "*"
	// 				} else {
	// 					multiplicity = "1"
	// 				}
	// 				fd.WriteString(fmt.Sprintf("%s 1--%s %s {label: \"%s - %s\"}\n", sKey, multiplicity, sourceRel.Target.Table, sourceRel.Id, targetRel.Id))
	// 			}
	// 		}
	// 	}

	// }

	return nil
}

func generateSingleRelation(relation Relation, entity Entity) (string, error) {
	var multiplicity string
	if relation.Type == "array" {
		multiplicity = "*"
	} else {
		multiplicity = "1"
	}
	return fmt.Sprintf("%s 1--%s %s {label: \"%s\"}\n", entity.Name, multiplicity, relation.Target.Table, relation.Id), nil
}

func generateSingleProperty(field Field, entity Entity) (string, error) {
	var pk string
	var fk string
	if field.Mandatory {
		pk = "*"
	} else {
		pk = ""
		fk = ""
		found := false
		// check if this field is a relation
		for _, relation := range entity.Relations {
			if found {
				break
			}
			for key := range relation.Target.Fields {
				if key == field.Id {
					fk = "+"
					found = true
					break
				}
			}
		}
	}
	return fmt.Sprintf("%s%s%s\n", pk, fk, field.Id), nil
}

func ParseMetadata(filePath string) Entity {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	var entity Entity
	err = json.NewDecoder(file).Decode(&entity)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return entity
}
