package main

import (
	"path/filepath"
	"io/ioutil"
	"os"
	"go/build"
	"fmt"
	"encoding/json"
	"strings"
	"text/template"
	"io"
)

func main()  {
	args := os.Args

	fmt.Println("Generating flogo metadata in ", args[1])

	generateGoMetadata(args[1])
}

func generateGoMetadata(dir string) error {
	//todo optimize metadata recreation to minimize compile times
	dependencies, err := ListDependencies(dir)

	if err != nil {
		return err
	}

	for _, dependency := range dependencies {

		fmt.Println("Generating flogo metadata for:", dependency.Ref)

		createMetadata(dependency)
	}

	return nil
}

func createMetadata(dependency *Dependency) error {

	var mdFilePath string
	var mdGoFilePath string
	var tplMetadata string

	switch dependency.ContribType {
	case ACTION:
		mdFilePath = filepath.Join(dependency.Dir, "action.json")
		mdGoFilePath = filepath.Join(dependency.Dir, "action_metadata.go")
		tplMetadata = tplMetadataGoFile
	case TRIGGER:
		mdFilePath = filepath.Join(dependency.Dir, "trigger.json")
		mdGoFilePath = filepath.Join(dependency.Dir, "trigger_metadata.go")
		tplMetadata = tplTriggerMetadataGoFile
	case ACTIVITY:
		mdFilePath = filepath.Join(dependency.Dir, "activity.json")
		mdGoFilePath = filepath.Join(dependency.Dir, "activity_metadata.go")
		tplMetadata = tplActivityMetadataGoFile
	default:
		return nil
	}

	raw, err := ioutil.ReadFile(mdFilePath)
	if err != nil {
		return err
	}

	info := &struct {
		Package      string
		MetadataJSON string
	}{
		Package:      filepath.Base(dependency.Dir),
		MetadataJSON: string(raw),
	}

	f, _ := os.Create(mdGoFilePath)
	RenderTemplate(f, tplMetadata, info)
	f.Close()

	return nil
}

var tplMetadataGoFile = `package {{.Package}}

var jsonMetadata = ` + "`{{.MetadataJSON}}`" + `

func getJsonMetadata() string {
	return jsonMetadata
}
`

var tplActivityMetadataGoFile = `package {{.Package}}

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = ` + "`{{.MetadataJSON}}`" + `

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
`

var tplTriggerMetadataGoFile = `package {{.Package}}

import (
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

var jsonMetadata = ` + "`{{.MetadataJSON}}`" + `

// init create & register trigger factory
func init() {
	md := trigger.NewMetadata(jsonMetadata)
	trigger.RegisterFactory(md.ID, NewFactory(md))
}
`

func ListDependencies(dir string) ([]*Dependency, error) {

	var cType ContribType

	// Get build context
	bc := build.Default
	currentGoPath := bc.GOPATH
	bc.GOPATH = dir
	//dir, _ := os.Getwd()
	//fmt.Println("gopath:", dir)
	ndir := "."
	//fmt.Println("dir:", ndir)

	defer func() { bc.GOPATH = currentGoPath }()
	pkgs, err := bc.ImportDir(ndir, 0)
	if err != nil {
		//fmt.Println("err:", err)
		return nil, err
	}
	//fmt.Println("pkgs:",pkgs)

	var deps []*Dependency
	// Get all imports
	for _, imp := range pkgs.Imports {

		//fmt.Println("imp:",imp)

		pkg, err := bc.Import(imp, ndir, build.FindOnly)
		if err != nil {
			//fmt.Println("import err:",err)
			// Ignore package
			continue
		}

		//fmt.Println("pkg:",pkg.Dir)

		if cType == 0 || cType == ACTION {
			filePath := filepath.Join(pkg.Dir, "action.json")
			// Check if it is an action
			info, err := os.Stat(filePath)
			if err == nil {
				desc, err := readDescriptor(filePath, info)
				if err == nil && desc.Type == "flogo:action" {
					deps = append(deps, &Dependency{ContribType: ACTION, Ref: imp, Dir: pkg.Dir})
				}
			}
		}
		if cType == 0 || cType == TRIGGER {
			filePath := filepath.Join(pkg.Dir, "trigger.json")
			// Check if it is a trigger
			info, err := os.Stat(filePath)
			if err == nil {
				desc, err := readDescriptor(filePath, info)
				if err == nil && desc.Type == "flogo:trigger" {
					deps = append(deps, &Dependency{ContribType: TRIGGER, Ref: imp, Dir: pkg.Dir})
				}
			}
		}
		if cType == 0 || cType == ACTIVITY {
			filePath := filepath.Join(pkg.Dir, "activity.json")
			// Check if it is an activity
			info, err := os.Stat(filePath)
			if err == nil {
				desc, err := readDescriptor(filePath, info)
				if err == nil && desc.Type == "flogo:activity" {
					deps = append(deps, &Dependency{ContribType: ACTIVITY, Ref: imp, Dir: pkg.Dir})
				}
			}
		}
	}
	return deps, nil
}

func readDescriptor(path string, info os.FileInfo) (*Descriptor, error) {

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error: " + err.Error())
		return nil, err
	}

	return ParseDescriptor(string(raw))
}

// ParseDescriptor parse a descriptor
func ParseDescriptor(descJson string) (*Descriptor, error) {
	descriptor := &Descriptor{}

	err := json.Unmarshal([]byte(descJson), descriptor)

	if err != nil {
		return nil, err
	}

	return descriptor, nil
}

type Descriptor struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Dependency struct {
	ContribType ContribType
	Ref         string
	Dir         string
}

func (d *Dependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ContribType string `json:"type"`
		Ref         string `json:"ref"`
	}{
		ContribType: d.ContribType.String(),
		Ref:         d.Ref,
	})
}

func (d *Dependency) UnmarshalJSON(data []byte) error {
	ser := &struct {
		ContribType string `json:"type"`
		Ref         string `json:"ref"`
	}{}

	if err := json.Unmarshal(data, ser); err != nil {
		return err
	}

	d.Ref = ser.Ref
	d.ContribType = ToContribType(ser.ContribType)

	return nil
}

type ContribType int

const (
	ACTION ContribType = 1 + iota
	TRIGGER
	ACTIVITY
)

var ctStr = [...]string{
	"all",
	"action",
	"trigger",
	"activity",
}

func (m ContribType) String() string { return ctStr[m] }

func ToContribType(name string) ContribType {
	switch name {
	case "action":
		return ACTION
	case "trigger":
		return TRIGGER
	case "activity":
		return ACTIVITY
	case "all":
		return 0
	}

	return -1
}

//RenderTemplate renders the specified template
func RenderTemplate(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
