package project

const ConfigFile = ".direktiv.yaml"
const DefaultCloudEventFiltersDirectory = ".direktiv.filters"
const DefaultNamespaceServicesDirectory = ".direktiv.services"

type DirektivProjectImport struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
	Ref  string `yaml:"yaml"`
}

type DirektivProjectSecret struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type DirektivProjectFileOverrides struct {
	Path     string `yaml:"path"`
	MIMEType string `yaml:"mimeType"`
	Type     string `yaml:"type"`
	Name     string `yaml:"name"`
}

type DirektivProjectConfig struct {
	Ignore            []string                       `yaml:"ignore"`
	Import            []DirektivProjectImport        `yaml:"import"`
	CloudEventFilters string                         `yaml:"cloudEventFilters"`
	NamespaceServices string                         `yaml:"namespaceServices"`
	Secrets           []DirektivProjectSecret        `yaml:"secrets"`
	FileOverrides     []DirektivProjectFileOverrides `yaml:"fileOverrides"`
}
