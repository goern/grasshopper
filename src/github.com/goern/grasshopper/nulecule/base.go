package nulecule

//GRASSHOPPERVERSION is Grasshopper version
const GRASSHOPPERVERSION = "0.0.2"

//NULECULESPECVERSION is the version of the Nulecule specification that Grasshopper is implementing
const NULECULESPECVERSION = "0.0.2"

//MAIN_FILE is the filename we are looking for within the Grasshopper based container
const MAIN_FILE = "Nulecule"

//ANSWERS_FILE is the filename we are looking for within the working context to get answers from (during install phase)
const ANSWERS_FILE = "answers.conf"

//ANSWERS_FILE_SAMPLE is the filename we write default answers to (during fetch phase)
const ANSWERS_FILE_SAMPLE = "answers.conf.sample"

//SUPPORTED_PROVIDERS lists all prividers that Grasshopper supports
const SUPPORTED_PROVIDERS = "none"

//DEFAULT_PROVIDER is the default provider to use
const DEFAULT_PROVIDER = "none"

//DEFAULT_NAMESPACE is the default namespace to be used with the provider
const DEFAULT_NAMESPACE = "default"

//LOCK_FILE location
const LOCK_FILE = "/run/lock/grasshopper.lock"

//HOST_DIR location
const HOST_DIR = "/host"

//Answers is a map of configuration parameters and the value to set for them
type Answers map[string]string

//Base contains a set of nulecule config properties
//It is set by the atomicapp subcommands
type Base struct {
	AnswersData        map[string]Answers
	answersDir         string
	MainfileData       *ContainerApplication
	targetPath         string
	Nodeps             bool
	DryRun             bool
	AppPath            string
	app                string
	WriteSampleAnswers bool
}

//ContainerApplication is a struct representation of the Nulecule specification file, see http://www.projectatomic.io/nulecule/spec/0.0.2/index.html
type ContainerApplication struct {
	ID           string
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
