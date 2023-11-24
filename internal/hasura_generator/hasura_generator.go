package hasura_generator

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	//"slices"
	//"strings"
	"github.com/go-resty/resty/v2"
)

func GenerateFromHasura(hasuraUrl string, adminSecret string, outputFile string, includeProperties bool) error {

//	var entities []Entity

//	var relationships = make(map[string][]*Relation)

	fd, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}
	defer fd.Close()


	url := fmt.Sprintf("%v/v1/query", hasuraUrl)

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-Hasura-Admin-Secret", adminSecret).
		SetBody(`{"type":"export_metadata","args":{}}`).
//		SetError(HasuraError{}).
		SetResult(HasuraMetadata{}).
		Post(url)

	if err != nil {
		log.Fatalf("Error getting metadata: %s", err)
	}

	printOutput(resp, err)


	log.Printf("Response: %v", resp)
	log.Printf("Body: %s", resp.Body())

	HasuraMetadata := resp.Result().(*HasuraMetadata)
	for _, source := range HasuraMetadata.Sources {
		if (source.Name == "data_access_manager") {
			continue
		}
		for _, table := range source.Tables {
			if table.Table.Schema != "public" {
				continue
			}
			fd.WriteString(fmt.Sprintf("[%s]\n", table.Table.Name))

			for _, orel := range table.ObjectRelationships {
				fd.WriteString(fmt.Sprintf("%s 1--1 %s {label: \"%s\"}\n", table.Table.Name, orel.Using.ManualConfiguration.RemoteTable.Name, orel.Name))
			}

			for _, arel := range table.ArrayRelationships {
				fd.WriteString(fmt.Sprintf("%s 1--* %s {label: \"%s\"}\n", table.Table.Name, arel.Using.ManualConfiguration.RemoteTable.Name, arel.Name))
			}
		}
	}
	// dir, err := os.Open(metadataDir)
	// if err != nil {
	// 	log.Fatalf("Error opening directory: %v", err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)
	// if err != nil {
	// 	log.Fatalf("Error reading directory: %v", err)
	// }

	// for _, fileName := range fileInfo {
	// 	if !strings.HasSuffix(fileName.Name(), ".json") {
	// 		continue
	// 	}

	// 	entity := parseMetadata(dir.Name() + "/" + fileName.Name())

	// 	entities = append(entities, entity)
	// }

	// for _, entity := range entities {
	// 	fd.WriteString(fmt.Sprintf("[%s]\n", entity.Name))
	// 	if includeProperties {
	// 		for _, field := range entity.Fields {
	// 			s, err := generateSingleProperty(field, entity)
	// 			if err != nil {
	// 				log.Fatalf("Error generating property: %v", err)
	// 			}
	// 			fd.WriteString(s)

	// 			if strings.HasSuffix(field.Id, "_lookup") {
	// 				name, _ := strings.CutSuffix(field.Id, "_lookup")
	// 				table := strings.Title(name)
	// 				log.Printf("Found lookup table %s", table)
	// 				// do we have a custom table in same namespace?
	// 				ns := strings.Split(entity.Name, "_")[0]
	// 				lookupTableName := ns + "_" + table
	// 				lookupTableIdx := slices.IndexFunc(entities, func(e Entity) bool {
	// 					return e.Name == lookupTableName || e.Name == lookupTableName+"s"
	// 				})
	// 				if lookupTableIdx != -1 {
	// 					lookupTable := entities[lookupTableIdx]
	// 					fd.WriteString(fmt.Sprintf("%s 1--1 %s {label: \"%s\"}\n", entity.Name, lookupTable.Name, field.Id))
	// 				} else {
	// 					fd.WriteString(fmt.Sprintf("%s 1--1 %s {label: \"%s\"}\n", entity.Name, "General_Lookups", field.Id))
	// 				}
	// 			}
	// 		}
	// 	}
	// 	fd.WriteString(fmt.Sprintf("\n"))

	// 	//relationships[entity.Name] = entity.Relations

	// 	for _, relation := range entity.Relations {
	// 		if relationships[entity.Name] == nil {
	// 			relationships[entity.Name] = make([]*Relation, 0)
	// 		}
	// 		relation.FromTable = entity.Name

	// 		relationships[entity.Name] = append(relationships[entity.Name], &relation)
	// 		if relationships[relation.Target.Table] == nil {
	// 			relationships[relation.Target.Table] = make([]*Relation, 0)
	// 		}
	// 		relationships[relation.Target.Table] = append(relationships[relation.Target.Table], &relation)
	// 		s, err := generateSingleRelation(relation, entity)
	// 		if err != nil {
	// 			log.Fatalf("Error generating relation: %v", err)
	// 		}
	// 		fd.WriteString(s)
	// 	}
	// 	fd.WriteString(fmt.Sprintf("\n"))
	// }

	// Skip self-references
	// do we have a reverse relationship?
	// rels, err := CollectRelations(relationships)
	// if err != nil {
	// 	log.Fatalf("Error collecting relations: %v", err)
	// }
	// for _, rel := range rels {
	// 	fd.WriteString(rel)
	// }

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

// func collectRelations(relationships map[string][]*Relation) ([]string, error) {
// 	var rels []string
// 	for sourceTable, relations := range relationships {
// 		log.Printf("\n\nChecking %s", sourceTable)
// 		slices.SortStableFunc(relations, func(a, b *Relation) int {
// 			return strings.Compare(a.Id, b.Id)
// 		})
// 		//relations = slices.Compact(relations)

// 		for _, relation := range relations {
// 			if relation.isHandled {
// 				continue
// 			}
// 			log.Printf(" - rel %s %v", relation.Id, relation)
// 			if sourceTable == relation.Target.Table {
// 				// Skip self-references
// 				continue
// 			}

// 			idx := slices.IndexFunc(relations, func(n *Relation) bool {
// 				return (n.Target.Table == relation.FromTable || relation.Target.Table == n.FromTable) &&
// 					relation.Id != n.Id
// 			})

// 			if idx != -1 {
// 				reverseRelation := relations[idx]
// 				var multiplicity string
// 				if relation.Type == "array" {
// 					multiplicity = "*"
// 				} else {
// 					multiplicity = "1"
// 				}

// 				var reverseMultiplicity string
// 				if reverseRelation.Type == "array" {
// 					reverseMultiplicity = "*"
// 				} else {
// 					reverseMultiplicity = "1"
// 				}

// 				relation.isHandled = true
// 				reverseRelation.isHandled = true
// 				log.Printf("Found reverse relationship %s (%s) between %s and %s", relation.Id, reverseRelation.Id, sourceTable, relation.Target.Table)
// 				rels = append(rels, fmt.Sprintf("%s %s--%s %s {label: \"%s - %s\"}\n", sourceTable, reverseMultiplicity, multiplicity, relation.Target.Table, relation.Id, reverseRelation.Id))
// 			} else {
// 				// no reverse relation found
// 				var multiplicity string
// 				if relation.Type == "array" {
// 					multiplicity = "*"
// 				} else {
// 					multiplicity = "1"
// 				}
// 				relation.isHandled = true
// 				log.Printf("Found single relationship %s between %s and %s", relation.Id, sourceTable, relation.Target.Table)

// 				rels = append(rels, fmt.Sprintf("%s 1--%s %s {label: \"%s\"}\n", sourceTable, multiplicity, relation.Target.Table, relation.Id))
// 			}
// 		}

// 	}
// 	return rels, nil
// }

// func generateSingleRelation(relation Relation, entity Entity) (string, error) {
// 	var multiplicity string
// 	if relation.Type == "array" {
// 		multiplicity = "*"
// 	} else {
// 		multiplicity = "1"
// 	}
// 	return fmt.Sprintf("%s 1--%s %s {label: \"%s\"}\n", entity.Name, multiplicity, relation.Target.Table, relation.Id), nil
// }

// func generateSingleProperty(field Field, entity Entity) (string, error) {
// 	var pk string
// 	var fk string
// 	if field.Mandatory {
// 		pk = "*"
// 	} else {
// 		pk = ""
// 		fk = ""
// 		found := false
// 		// check if this field is a relation
// 		for _, relation := range entity.Relations {
// 			if found {
// 				break
// 			}
// 			for key := range relation.Target.Fields {
// 				if key == field.Id {
// 					fk = "+"
// 					found = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return fmt.Sprintf("%s%s%s\n", pk, fk, field.Id), nil
// }

// func parseMetadata(filePath string) Entity {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatalf("Error opening file: %v", err)
// 	}
// 	var entity Entity
// 	err = json.NewDecoder(file).Decode(&entity)
// 	if err != nil {
// 		log.Fatalf("Error decoding JSON: %v", err)
// 	}
// 	return entity
// }
func printOutput(resp *resty.Response, err error) {
	fmt.Println(resp, err)
}
