package wazuh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/iancoleman/strcase"
)

type Dict map[string]interface{}

var Types []string = make([]string, 0)

func mapType(name string, spec map[string]interface{}) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("type %s struct {\n", strcase.ToCamel(name)))

	for jKey, v := range spec {
		spec := v.(map[string]interface{})
		var gKey string
		if strings.HasPrefix(jKey, "@") {
			gKey = strcase.ToCamel(jKey[1:]) + "_"
		} else {
			gKey = strcase.ToCamel(jKey)
		}
		if jKey == "id" {
			gKey = "ID"
		}
		if jKey == "url" {
			gKey = "URL"
		}
		if jKey == "ip" {
			gKey = "IP"
		}

		if jType, ok := spec["type"]; ok {
			var gType = ""

			if name == "rule" {
				if jType == "pci_dss" || jType == "tsc" || jType == "groups" || jType == "gdpr" || jType == "gpg13" || jType == "hipaa" || jType == "nist_800_53" {
					gType = "[]string"
				}
			}
			if name == "vulnerability" {
				if jType == "references" || jType == "bugzilla_references" {
					gType = "[]string"
				}
			}
			if name == "mitre" {
				if jType == "id" || jType == "technique" || jType == "tactic" {
					gType = "[]string"
				}
			}

			if gType == "" {

				switch t := jType; t {
				case "ip":
					gType = "*string"
				case "text":
					gType = "*string"
				case "keyword":
					gType = "*string"
				case "date":
					gType = "*int64"
				case "integer":
					gType = "*int64"
				case "long":
					gType = "*uint64"
				case "boolean":
					gType = "*bool"
				case "double":
					gType = "*float64"
				case "geo_point":
					gType = "*GeoPoint"
				default:
					gType = fmt.Sprintf("<%s>", jType)
				}
			}
			sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`\n", gKey, gType, jKey))
		} else if jSubProps, ok := spec["properties"]; ok {
			if jKey == "vector" || jKey == "process" || jKey == "check" || jKey == "audit" {
				jKey = name + "_" + jKey
			}
			mapType(jKey, jSubProps.(map[string]interface{}))
			sb.WriteString(fmt.Sprintf("\t%s *%s `json:\"%s,omitempty\"`\n", gKey, strcase.ToCamel(jKey), jKey))
		}
	}
	sb.WriteString("}\n")

	Types = append(Types, sb.String())
	return sb.String()
}

func TestImport(t *testing.T) {
	rawJson, err := ioutil.ReadFile("index-pattern.json")
	if err != nil {
		t.Fatal(err)
	}

	var dict Dict
	err = json.Unmarshal(rawJson, &dict)
	if err != nil {
		t.Fatal(err)
	}

	rawMapping := dict["mappings"].(map[string]interface{})
	propertieDefs := rawMapping["properties"].(map[string]interface{})
	mapType("Alert", propertieDefs)

	for _, t := range Types {
		fmt.Printf("%v\n", t)
	}

}
