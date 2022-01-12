package utils

import (
	"bytes"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/orangedeng/rancher-api-doc-generator/pkg/types"
	md "github.com/orangedeng/rancher-api-doc-generator/pkg/types"
	"github.com/rancher/norman/clientbase"
	normantypes "github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	title    string
	endpoint string
	ak, sk   string
)

func GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:        "title,t",
			Destination: &title,
			EnvVar:      "TITLE",
			Value:       "Rancher Norman Stype API Docs",
			Usage:       "the title of the API docs",
		},
		cli.StringFlag{
			Name:        "endpoint,e",
			Required:    true,
			Usage:       "the API endpoint for generating API docs",
			Destination: &endpoint,
			EnvVar:      "ENDPOINT",
		},
		cli.StringFlag{
			Name:        "access-key",
			Usage:       "the access key for accessing API endpoint",
			EnvVar:      "ACCESS_KEY",
			Destination: &ak,
		},
		cli.StringFlag{
			Name:        "secret-key",
			Usage:       "the secret key for accessing API endpoint",
			EnvVar:      "SECRET_KEY",
			Destination: &sk,
		},
	}
}

func GenerateDocs() (*md.MD, error) {
	client, err := clientbase.NewAPIClient(&clientbase.ClientOpts{
		URL:       endpoint,
		AccessKey: ak,
		SecretKey: sk,
		Insecure:  true,
	})
	if err != nil {
		return nil, err
	}
	rtn := md.NewMarkDown()

	if err := GenerateTitle(rtn); err != nil {
		return rtn, err
	}

	if err := GenerateSchemaList(client, rtn); err != nil {
		return rtn, err
	}

	if err := GenerateTypeList(client, rtn); err != nil {
		return rtn, err
	}

	if err := GenerateAPIList(client, rtn); err != nil {
		return rtn, err
	}

	return rtn, nil
}

func GenerateTitle(file *md.MD) error {
	file.Append(md.H1Heading(title))
	file.Append(md.Paragraph("This is the API docs for ", md.Link{Anchor: endpoint, Link: endpoint}, "."))
	return nil
}

func GenerateSchemaList(client clientbase.APIBaseClient, file *md.MD) error {
	file.Append(md.H2Heading("Schema List"))
	file.Append(md.Paragraph("Followings are the API links in the level."))
	list := md.NewList(false)
	words := []string{}
	for typeName, schema := range client.Types {
		if _, ok := schema.Links["collection"]; !ok {
			continue
		}
		words = append(words, md.NewWords(
			md.Link{Anchor: typeName, Link: "#schema-" + strings.ToLower(typeName) + "-api-list"}, " ",
			md.Link{Anchor: "example", Link: schema.Links["collection"]},
		).String())
	}
	sort.Strings(words)
	for _, line := range words {
		list = list.Append(md.SimpleWord(line), nil)
	}
	file.Append(list)
	return nil
}

func GenerateTypeList(client clientbase.APIBaseClient, file *md.MD) error {
	file.Append(md.H2Heading("API Type List"))
	file.Append(md.Paragraph("Following are the schema types in this endpoint."))
	var schemaNames []string
	for schemaName := range client.Types {
		schemaNames = append(schemaNames, schemaName)
	}
	sort.Strings(schemaNames)
	// for _, typeName := range schemaNames {
	// 	schema := client.Types[typeName]
	// 	tables := getSchemaTables(typeName, schema, client.Types)
	// 	file.Append(tables.standard)
	// }

	return nil
}

func splitType(t string) (string, string) {
	prefixIndex := strings.IndexRune(t, '[')
	if prefixIndex == -1 {
		return "", t
	}
	return t[:prefixIndex], t[prefixIndex+1 : len(t)-1]
}

func getTypeWord(t string) md.Words {
	prefix, content := splitType(t)
	switch prefix {
	case "array", "map":
		return md.NewWords(prefix + "[" + getTypeWord(content).String() + "]")
	case "reference":
		return md.Link{Anchor: content, Link: "#" + content}
	}
	return md.SimpleWord(t)
}

func getFieldLengthWords(schema normantypes.Field) md.Words {
	var max, min int64
	if schema.Max != nil {
		max = *schema.Max
	}
	if schema.MaxLength != nil {
		max = *schema.MaxLength
	}
	if schema.Min != nil {
		min = *schema.Min
	}
	if schema.MinLength != nil {
		min = *schema.MinLength
	}
	if max == 0 && min == 0 {
		return md.SimpleWord("-")
	}
	var Max, Min string
	if min == 0 {
		Min = "\\*"
	} else {
		Min = fmt.Sprintf("%d", min)
	}
	if max == 0 {
		Max = "\\*"
	} else {
		Max = fmt.Sprintf("%d", max)
	}
	return md.SimpleWord(fmt.Sprintf("%s - %s", Min, Max))
}

func getFieldCommentWords(field normantypes.Field) md.Words {
	buffer := bytes.NewBuffer([]byte{})

	for _, kv := range []map[string]string{
		{"description:": field.Description},
		{"enum options:": strings.Join(field.Options, "/")},
		{"valid chars:": field.ValidChars},
		{"invalid chars:": field.InvalidChars},
	} {
		for k, v := range kv {
			if v != "" {
				buffer.WriteString(fmt.Sprintf("%s %s; ", k, v))
			}
		}
	}

	return md.SimpleWord(buffer.String())
}

func GenerateAPIList(client clientbase.APIBaseClient, file *md.MD) error {
	// title
	file.Append(md.H2Heading("API List"))
	file.Append(md.Paragraph("This is the API List of this endppint."))
	var schemas []string
	for typeName, schema := range client.Types {
		if _, ok := schema.Links["collection"]; !ok {
			continue
		}
		schemas = append(schemas, typeName)
	}
	sort.Strings(schemas)
	for _, name := range schemas {
		schema := client.Types[name]
		// Type Title
		file.Append(md.H3Heading("Schema", name, "API List"))
		for _, method := range schema.CollectionMethods {
			file.Append(generateMethod(method, true, schema, client.Types))
		}
		for _, method := range schema.ResourceMethods {
			file.Append(generateMethod(method, false, schema, client.Types))
		}
		for name, action := range schema.ResourceActions {
			file.Append(generateAction(false, name, action, schema, client.Types))
		}
	}
	return nil
}

func generateMethod(method string, isPlural bool, schema normantypes.Schema, schemas map[string]normantypes.Schema) md.MultipleLine {
	//generating API header
	var uri string
	collectionURL, err := url.Parse(schema.Links["collection"])
	if err != nil {
		logrus.Panicf("failed to parse collection link from schema %s", schema.ID)
	}
	uri = collectionURL.RequestURI()
	if !isPlural {
		uri += "/" + types.Italic("resource-id").String()
	}
	rtn := append([]interface{}{}, md.H4Heading(method, uri))

	//generating API input and output
	switch strings.ToUpper(method) {
	case "GET", "DELETE":
		rtn = append(rtn, "\n")
		target := md.NewList(false).Append(md.NewWords("input: null"), nil)
		if isPlural {
			target = target.Append(md.NewWords("output: collection of ", md.Link{Anchor: schema.ID, Link: toTypeLink(schema.ID)}), nil)
		} else {
			target = target.Append(md.NewWords("output: ", md.Link{Anchor: schema.ID, Link: toTypeLink(schema.ID)}), nil)
		}
		rtn = append(rtn, target)
	case "POST", "PUT":
		rtn = append(rtn, "\n")
		target := md.NewList(false).
			Append(md.NewWords("input: ", md.Link{Anchor: schema.ID, Link: toTypeLink(schema.ID)}), nil).
			Append(md.NewWords("output: ", md.Link{Anchor: schema.ID, Link: toTypeLink(schema.ID)}), nil)
		rtn = append(rtn, target)
	}
	return md.Paragraph(rtn...)
}

func generateAction(isPlural bool, actionName string, action normantypes.Action, schema normantypes.Schema, schemas map[string]normantypes.Schema) md.MultipleLine {
	//generating API header
	var uri string
	collectionURL, err := url.Parse(schema.Links["collection"])
	if err != nil {
		logrus.Panicf("failed to parse collection link from schema %s", schema.ID)
	}
	uri = collectionURL.RequestURI()
	if !isPlural {
		uri += "/" + types.Italic("resource-id").String() + "?action=" + actionName
	} else {
		uri += "?action=" + actionName
	}
	rtn := []interface{}{
		types.H4Heading(fmt.Sprintf("Action %s of schema object %s, POST %s", actionName, schema.ID, uri)),
		"\n",
	}

	//generating input and output for action

	target := md.NewList(false)
	if action.Input != "" {
		target = target.Append(md.NewWords("input: ", md.Link{Anchor: action.Input, Link: toTypeLink(action.Input)}), nil)
	} else {
		target = target.Append(md.SimpleWord("input: null"), nil)
	}
	if action.Output != "" {
		target = target.Append(md.NewWords("output: ", md.Link{Anchor: action.Output, Link: toTypeLink(action.Output)}), nil)
	} else {
		target = target.Append(md.SimpleWord("output: null"), nil)
	}
	rtn = append(rtn, target)
	return md.Paragraph(rtn...)
}

func toTypeLink(schemaName string) string {
	return fmt.Sprintf("#%s", strings.ToLower(schemaName))
}
