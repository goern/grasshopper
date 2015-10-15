// Package nulecule will provide some constants required for Grasshopper
// and all required data structures to run a Nulecule.
package nulecule

const (
	//GrasshopperVersion is Grasshopper version
	GrasshopperVersion = "0.0.3"

	//NuleculeVersion is the version of the Nulecule specification that Grasshopper is implementing
	NuleculeVersion = "0.0.2"

	//NulefileManifestFile is the filename we are looking for within the Grasshopper based container
	NulefileManifestFile = "Nulecule"

	//AnswersFile is the filename we are looking for within the working context to get answers from (during install phase)
	AnswersFile = "answers.conf"

	//AnswersFileSample is the filename we write default answers to (during fetch phase)
	AnswersFileSample = "answers.conf.sample"

	//SupportedProviders lists all prividers that Grasshopper supports
	SupportedProviders = "none"

	//DefaultProvider is the default provider to use
	DefaultProvider = "none"

	//DefaultNamespace is the default namespace to be used with the provider
	DefaultNamespace = "default"

	//GrasshopperLockFile location
	GrasshopperLockFile = "/run/lock/grasshopper.lock"

	//HostDir location
	HostDir = "/host"
)

//Answers is a map of configuration parameters and their value
type Answers map[string]string

//Base contains a set of nulecule config properties
type Base struct {
	AnswersData            map[string]Answers
	ContainerApplication   *ContainerApplication
	WriteAnswersFileSample bool
}

//ContainerApplication is a struct representating the Nulecule file, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html
type ContainerApplication struct {
	AppID        string
	Specversion  string
	Metadata     *Metadata
	Graph        []Component
	Requirements []Requirement
}

//Metadata arbitrary_data is not supported by Grasshopper, represents a http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#metadataObject
type Metadata struct {
	Name        string
	AppVersion  string
	Description string
	License     License
}

//License represents a http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#licenseObject
type License struct {
	Name string
	URL  string
}

//Param represents the Component parameters, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#paramsObject
type Param struct {
	Name        string
	Description string
	Constraints []Constraint
	Default     string
	Hidden      bool
}

//Constraint is a struct representing a constaint for a parameter object, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#constraintObject
type Constraint struct {
	AllowedPattern string
	Description    string
}

//ArtifactEntry is a source control repository struct used to specify an artifact, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#artifactsObject
type ArtifactEntry struct {
	Path string
	Repo SrcControlRepo
}

//SrcControlRepo is a Source Control Repository Object for artifact sources, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#repositoryObject
type SrcControlRepo struct {
	Source string
	Path   string
	Type   string
	Branch string
	Tag    string
}

//Component represents a graph item of the Nulecule file, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#graphObject
type Component struct {
	Name      string
	Source    string
	Params    []Param
	Artifacts map[string][]ArtifactEntry
}

//Requirement is a list of requirements of the Container Application, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html#storageRequirementsObject, Grasshopper only supports Storage Requirement
type Requirement struct {
	Name       string
	AccessMode string
	Size       int
}
