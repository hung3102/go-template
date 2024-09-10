package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	flag "github.com/spf13/pflag"
	"golang.org/x/xerrors"
	"gopkg.in/yaml.v3"
)

type execParam struct {
	YamlPath, SchemasDir, ExamplesDir string
}

func getComponentList(dir, yamlPath string) (map[string]*openapi3.SchemaRef, error) {
	// key: schemaのYAMLパス, value: Objectの名前
	targetList := make(map[string]string)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return xerrors.Errorf("errors is not nil: %w", err)
		}
		if d.IsDir() {
			return nil
		}
		ap, err := filepath.Rel(dir, path)
		if err != nil {
			return xerrors.Errorf("failed to rel path: %w (%s : %s)", err, dir, path)
		}

		plu := pluralize.NewClient()
		var cs []string
		for _, p := range strings.Split(ap, "/") {
			cs = append(cs, plu.Singular(p))
		}

		cp := strings.ReplaceAll(strings.Join(cs, "/"), "/", "_")
		sp := cp[:len(cp)-len(filepath.Ext(cp))]

		sn := strings.NewReplacer("Get", "", "Post", "", "Put", "", "Delete", "").Replace(strcase.ToCamel(sp))
		yd := filepath.Dir(yamlPath)
		op, err := filepath.Rel(yd, path)
		if err != nil {
			return xerrors.Errorf("failed to rel path: %w (%s: %s)", err, yd, path)
		}
		targetList[op] = sn

		return nil
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to filepath.WalkDir: %w", err)
	}

	ss := make(map[string]*openapi3.SchemaRef)
	for path, name := range targetList {
		ss[name] = &openapi3.SchemaRef{
			Ref: fmt.Sprintf("./%s", path),
		}
	}

	return ss, nil
}

func exec(param execParam) error {
	var spec map[string]any
	yamlPath := param.YamlPath
	schemasDir := param.SchemasDir
	examplesDir := param.ExamplesDir

	specBin, err := os.ReadFile(yamlPath)
	if err != nil {
		return xerrors.Errorf("failed to read open api file: %w", err)
	}

	err = yaml.Unmarshal(specBin, &spec)
	if err != nil {
		return xerrors.Errorf("failed to unmarshal yaml: %w", err)
	}

	// components 一覧を作る
	{
		ss, err := getComponentList(schemasDir, yamlPath)
		if err != nil {
			return xerrors.Errorf("failed to listed component (components): %w", err)
		}

		spec["components"].(map[string]any)["schemas"] = ss
	}

	// examples 一覧を作る
	{
		es, err := getComponentList(examplesDir, yamlPath)
		if err != nil {
			return xerrors.Errorf("failed to listed component (examples): %w", err)
		}

		spec["components"].(map[string]any)["examples"] = es
	}

	y, err := yaml.Marshal(spec)
	if err != nil {
		return xerrors.Errorf("failed to marshal yaml: %w", err)
	}

	fmt.Println(string(y))

	return nil
}

func main() {
	var (
		openapiYaml = flag.String("openapi", "", "./openapi/openapi.yaml")
		schemasDir  = flag.String("schemas-dir", "", "./openapi/schemas/")
		examplesDir = flag.String("examples-dir", "", "./openapi/examples/")
	)
	flag.Parse()

	err := exec(execParam{
		YamlPath:    *openapiYaml,
		SchemasDir:  *schemasDir,
		ExamplesDir: *examplesDir,
	})
	if err != nil {
		log.Fatalf("failed to exec: %v", err)
	}
}
