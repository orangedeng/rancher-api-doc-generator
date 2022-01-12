package utils

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	md "github.com/orangedeng/rancher-api-doc-generator/pkg/types"
	normantypes "github.com/rancher/norman/types"
	"github.com/stretchr/testify/assert"
)

func TestIsReference(t *testing.T) {
	type testcase struct {
		name        string
		source      string
		isReference bool
		realType    string
	}

	for _, c := range []testcase{
		{
			name:        "array wrap reference type",
			source:      "array[reference[abcd]]",
			isReference: true,
			realType:    "abcd",
		},
		{
			name:        "raw reference type",
			source:      "reference[abcd]",
			isReference: true,
			realType:    "abcd",
		},
		{
			name:        "array type",
			source:      "array[string]",
			isReference: false,
			realType:    "string",
		},
		{
			name:        "raw type",
			source:      "string",
			isReference: false,
			realType:    "string",
		},
	} {
		isRef, real := isReference(c.source)
		assert.Equalf(t, c.isReference, isRef, "test case %s failed, isReference return is not equal", c.name)
		assert.Equalf(t, c.realType, real, "test case %s failed, realType return is not equal", c.name)
	}
}

func TestGetSchemaTable(t *testing.T) {
	data, err := ioutil.ReadFile("./schemas.json")
	if err != nil {
		t.Fatal(err)
	}
	schemas := map[string]normantypes.Schema{}
	if err := json.Unmarshal(data, &schemas); err != nil {
		t.Fatal(err)
	}

	ctx := map[string]*typeTable{}
	type testcase struct {
		name       string
		schemaName string
		create     *md.Table
		update     *md.Table
		read       *md.Table
	}
	for _, c := range []testcase{} {
		getSchemaTables(ctx, c.schemaName, schemas)
		table := ctx[c.schemaName]
		assert.NotNil(t, table, "test case %s failed, failed to get schema table", c.name)
		assert.Equalf(t, c.create, table.create, "test case %s failed, create table not match", c.name)
		assert.Equalf(t, c.update, table.update, "test case %s failed, update table not match", c.name)
		assert.Equalf(t, c.read, table.read, "test case %s failed, read table not match", c.name)
	}
}
