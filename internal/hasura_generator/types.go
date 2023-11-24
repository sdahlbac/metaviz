package hasura_generator

type HasuraError struct {}

type HasuraMetadata struct {
	Version int `json:"version"`
	Sources []struct {
		Name   string `json:"name"`
		Kind   string `json:"kind"`
		Tables []struct {
			Table struct {
				Name   string `json:"name"`
				Schema string `json:"schema"`
			} `json:"table"`
			ObjectRelationships []struct {
				Name  string `json:"name"`
				Using struct {
					ForeignKeyConstraintOn string `json:"foreign_key_constraint_on"`
					ManualConfiguration struct {
						RemoteTable struct {
							Name   string `json:"name"`
						} `json:"remote_table"`
					} `json:"manual_configuration"`
				} `json:"using"`
			} `json:"object_relationships,omitempty"`
			// InsertPermissions []struct {
			// 	Role       string `json:"role"`
			// 	Permission struct {
			// 		Check struct {
			// 		} `json:"check"`
			// 		Columns string `json:"columns"`
			// 	} `json:"permission"`
			// } `json:"insert_permissions"`
			// SelectPermissions []struct {
			// 	Role       string `json:"role"`
			// 	Permission struct {
			// 		Columns string `json:"columns"`
			// 		Filter  struct {
			// 		} `json:"filter"`
			// 		AllowAggregations bool `json:"allow_aggregations"`
			// 	} `json:"permission"`
			// } `json:"select_permissions"`
			// UpdatePermissions []struct {
			// 	Role       string `json:"role"`
			// 	Permission struct {
			// 		Columns string `json:"columns"`
			// 		Filter  struct {
			// 		} `json:"filter"`
			// 		Check any `json:"check"`
			// 	} `json:"permission"`
			// } `json:"update_permissions"`
			// DeletePermissions []struct {
			// 	Role       string `json:"role"`
			// 	Permission struct {
			// 		Filter struct {
			// 		} `json:"filter"`
			// 	} `json:"permission"`
			// } `json:"delete_permissions"`
			ArrayRelationships []struct {
				Name  string `json:"name"`
				Using struct {
					ForeignKeyConstraintOn struct {
						Column string `json:"column"`
						Table  struct {
							Name   string `json:"name"`
							Schema string `json:"schema"`
						} `json:"table"`
					} `json:"foreign_key_constraint_on"`
					ManualConfiguration struct {
						RemoteTable struct {
							Name   string `json:"name"`
						} `json:"remote_table"`
					} `json:"manual_configuration"`
				} `json:"using"`
			} `json:"array_relationships,omitempty"`
		} `json:"tables"`
		Configuration struct {
			ConnectionInfo struct {
				DatabaseURL struct {
					FromEnv string `json:"from_env"`
				} `json:"database_url"`
				IsolationLevel        string `json:"isolation_level"`
				UsePreparedStatements bool   `json:"use_prepared_statements"`
			} `json:"connection_info"`
		} `json:"configuration"`
		Functions []struct {
			Function struct {
				Name   string `json:"name"`
				Schema string `json:"schema"`
			} `json:"function"`
		} `json:"functions,omitempty"`
	} `json:"sources"`
	Actions []struct {
		Name       string `json:"name"`
		Definition struct {
			Handler    string `json:"handler"`
			OutputType string `json:"output_type"`
			Type       string `json:"type"`
			Kind       string `json:"kind"`
		} `json:"definition,omitempty"`
		Comment     string `json:"comment"`
		Permissions []struct {
			Role string `json:"role"`
		} `json:"permissions"`
		Definition0 struct {
			Handler    string `json:"handler"`
			OutputType string `json:"output_type"`
			Arguments  []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"arguments"`
			Type string `json:"type"`
		} `json:"definition,omitempty"`
	} `json:"actions"`
	CustomTypes struct {
		InputObjects []struct {
			Name   string `json:"name"`
			Fields []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"fields"`
		} `json:"input_objects"`
		Objects []struct {
			Name   string `json:"name"`
			Fields []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"fields"`
		} `json:"objects"`
	} `json:"custom_types"`
}

// type HasuraMetadata struct {
// 	Version int `json:"version"`
// 	Sources []Source `json:"sources"`
// }

// type Source struct {
// 	Name string `json:"name"`
// 	Kind string `json:"kind"`
// 	Tables []Table `json:"tables"`
// }

// type Table struct {
// 	Schema string `json:"schema"`
// 	Name string `json:"name"`
// 	ObjectRelationships []ObjectRelationship `json:"object_relationships"`
// 	ArrayRelationships []ArrayRelationship `json:"array_relationships"`
// 	ComputedFields []ComputedField `json:"computed_fields"`
// }

// type ObjectRelationship struct {
// 	Name string `json:"name"`
// 	Using Using `json:"using"`
// }

// type Using struct {
// 	ManualConfiguration ManualConfiguration `json:"manual_configuration"`
// }

// type Entity struct {
// 	Name        string
// 	Description string
// 	Fields      []Field
// 	Relations   []Relation
// }

// type Field struct {
// 	Id          string
// 	Name        string
// 	Type        string
// 	Mandatory   bool
// 	Description string
// }

// type Relation struct {
// 	Type      string
// 	Id        string
// 	Name      string
// 	Target    Target
// 	isHandled bool
// 	FromTable string // Used to track which table the relation was found in
// }

// type Target struct {
// 	Table  string
// 	Fields map[string]interface{}
// }
