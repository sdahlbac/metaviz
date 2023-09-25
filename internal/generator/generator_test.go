package generator

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateSinglePropertyPK(t *testing.T) {
	field := Field{
		Id:        "test_field",
		Mandatory: true,
	}
	entity := Entity{
		Name:      "test_entity",
		Relations: []Relation{},
	}

	expected := "*test_field\n"
	actual, err := generateSingleProperty(field, entity)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}
func TestGenerateSinglePropertyVanilla(t *testing.T) {
	field := Field{
		Id:        "test_field",
		Mandatory: false,
	}
	entity := Entity{
		Name: "test_entity",
		Relations: []Relation{
			{
				Target: Target{
					Table: "test_target",
					Fields: map[string]interface{}{
						"test_field2": "Id",
					},
				},
			},
		},
	}

	expected := "test_field\n"
	actual, err := generateSingleProperty(field, entity)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestGenerateSinglePropertyFK(t *testing.T) {
	field := Field{
		Id:        "test_field",
		Mandatory: false,
	}
	entity := Entity{
		Name: "test_entity",
		Relations: []Relation{
			{
				Target: Target{
					Table: "test_target",
					Fields: map[string]interface{}{
						"test_field": "Id",
					},
				},
			},
		},
	}

	expected := "+test_field\n"
	actual, err := generateSingleProperty(field, entity)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestParseMetadata(t *testing.T) {
	// Create a temporary file with some test data
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up
	data := []byte(`{"name": "test_entity"}`)
	if _, err := tmpfile.Write(data); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test the function
	entity := ParseMetadata(tmpfile.Name())
	if entity.Name != "test_entity" {
		t.Errorf("Expected entity name to be %q, but got %q", "test_entity", entity.Name)
	}
}
