package utils

import (
	"fmt"
	"sort"

	md "github.com/orangedeng/rancher-api-doc-generator/pkg/types"
	normantypes "github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

type typeTable struct {
	schema   *normantypes.Schema
	standard *md.Table
	create   *md.Table
	read     *md.Table
	update   *md.Table
}

func getSchemaTables(ctx map[string]*typeTable, typeName string, schemas map[string]normantypes.Schema) {
	schema := schemas[typeName]
	standard := md.NewTable(
		md.TablesHeader{Title: "attribute name"},
		md.TablesHeader{Title: "type"},
		md.TablesHeader{Title: "C/R/U"},
		md.TablesHeader{Title: "required"},
		md.TablesHeader{Title: "nullable"},
		md.TablesHeader{Title: "length/range"},
		md.TablesHeader{Title: "comments"},
	)
	create := md.NewTable(
		md.TablesHeader{Title: "attribute name"},
		md.TablesHeader{Title: "type"},
		md.TablesHeader{Title: "required"},
		md.TablesHeader{Title: "nullable"},
		md.TablesHeader{Title: "length/range"},
		md.TablesHeader{Title: "comments"},
	)
	update := md.NewTable(
		md.TablesHeader{Title: "attribute name"},
		md.TablesHeader{Title: "type"},
		md.TablesHeader{Title: "required"},
		md.TablesHeader{Title: "nullable"},
		md.TablesHeader{Title: "length/range"},
		md.TablesHeader{Title: "comments"},
	)
	read := md.NewTable(
		md.TablesHeader{Title: "attribute name"},
		md.TablesHeader{Title: "type"},
		md.TablesHeader{Title: "length/range"},
		md.TablesHeader{Title: "comments"},
	)
	fieldNames := []string{}
	for name := range schema.ResourceFields {
		fieldNames = append(fieldNames, name)
	}
	sort.Strings(fieldNames)
	for _, name := range fieldNames {
		field := schema.ResourceFields[name]
		standard = standard.NewRow().
			Append(md.SimpleWord(name)).
			Append(getTypeWord(field.Type)).
			Append(getCRUWords(field)).
			Append(md.SimpleWord(fmt.Sprintf("%v", field.Required))).
			Append(md.SimpleWord(fmt.Sprintf("%v", field.Nullable))).
			Append(getFieldLengthWords(field)).
			Append(getFieldCommentWords(field)).
			Table()

		isReference, fieldRealType := isReference(field.Type)

		if isReference || canCreate(field) {
			create = create.NewRow().
				Append(md.SimpleWord(name)).
				Append(getTypeWord(field.Type)).
				Append(md.SimpleWord(fmt.Sprintf("%v", field.Required))).
				Append(md.SimpleWord(fmt.Sprintf("%v", field.Nullable))).
				Append(getFieldLengthWords(field)).
				Append(getFieldCommentWords(field)).
				Table()
		}

		if isReference || canUpdate(field) {
			update = update.NewRow().
				Append(md.SimpleWord(name)).
				Append(getTypeWord(field.Type)).
				Append(md.SimpleWord(fmt.Sprintf("%v", field.Required))).
				Append(md.SimpleWord(fmt.Sprintf("%v", field.Nullable))).
				Append(getFieldLengthWords(field)).
				Append(getFieldCommentWords(field)).
				Table()
		}

		if isReference || canRead(field) {
			read = read.NewRow().
				Append(md.SimpleWord(name)).
				Append(getTypeWord(field.Type)).
				Append(getFieldLengthWords(field)).
				Append(getFieldCommentWords(field)).
				Table()
		}

		if isReference {
			subtable, ok := ctx[fieldRealType]
			if !ok {
				getSchemaTables(ctx, fieldRealType, schemas)
				subtable = ctx[fieldRealType]
				if subtable == nil {
					logrus.Panicf("failed to ensure schema table")
				}
			}
			for _, row := range subtable.create.CloneRows() {
				currentName := name + create.GetRow(0).GetColumn(0).String()
				row.SetColumn(0, md.SimpleWord(currentName))
				create.Append(row)
			}
			for _, row := range subtable.update.CloneRows() {
				currentName := name + create.GetRow(0).GetColumn(0).String()
				row.SetColumn(0, md.SimpleWord(currentName))
				update.Append(row)
			}
			for _, row := range subtable.read.CloneRows() {
				currentName := name + create.GetRow(0).GetColumn(0).String()
				row.SetColumn(0, md.SimpleWord(currentName))
				read.Append(row)
			}
		}
	}

	ctx[typeName] = &typeTable{
		schema:   &schema,
		standard: standard,
		create:   create,
		update:   update,
		read:     read,
	}
}

func isReference(t string) (bool, string) {
	var currentPrefix string
	realType := t
	for prefix, fieldRealType := splitType(realType); prefix != ""; prefix, fieldRealType = splitType(fieldRealType) {
		currentPrefix = prefix
		realType = fieldRealType
	}
	return currentPrefix == "reference", realType
}

func getCRUWords(field normantypes.Field) md.Words {
	c, r, u := "-", "R", "-"
	if canCreate(field) {
		c = "C"
	}
	if canUpdate(field) {
		u = "U"
	}
	if !canRead(field) {
		r = "-"
	}
	return md.SimpleWord(fmt.Sprintf("%s/%s/%s", c, r, u))
}

func canCreate(field normantypes.Field) bool {
	return field.Create
}

func canUpdate(field normantypes.Field) bool {
	return field.Update
}

func canRead(field normantypes.Field) bool {
	return !field.WriteOnly
}
