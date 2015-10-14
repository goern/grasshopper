package nulecule

const GRASSHOPPERVERSION = "0.0.2"
const NULECULESPECVERSION = "0.0.2"

const MAIN_FILE = "Nulecule"
const ANSWERS_FILE = "answers.conf"
const ANSWERS_FILE_SAMPLE = "answers.conf.sample"

const SUPPORTED_PROVIDERS = "none"
const DEFAULT_PROVIDER = "openshift"
const DEFAULT_NAMESPACE = "default"

const LOCK_FILE = "/run/lock/atomicapp.lock"
const HOST_DIR = "/host"

//Answers is a map of configuration parameters and the value to set for them
type Answers map[string]string

//Base contains a set of nulecule config properties
//It is set by the atomicapp subcommands
type Base struct {
	AnswersData        map[string]Answers
	answersDir         string
	MainfileData       *Mainfile
	targetPath         string
	Nodeps             bool
	DryRun             bool
	AppPath            string
	app                string
	WriteSampleAnswers bool
}

//Mainfile is a struct representation of the Nulecule specification file
type Mainfile struct {
	Specversion string
	ID          string
	Graph       []Component
}

//Param represents the Component parameters
type Param struct {
	Name        string
	Description string
	Constraints []Constraint
	Default     string
	Hidden      bool
	AskedFor    bool
}

//Constraint is a struct representing a constaint for a parameter object
type Constraint struct {
	AllowedPattern string `json:"allowed_pattern",yaml:"allowed_pattern"`
	Description    string
}

//ArtifactEntry is a source control repository struct used to specify an artifact
type ArtifactEntry struct {
	Path string
	Repo SrcControlRepo
}

//SrcControlRepo is a ...
type SrcControlRepo struct {
	Inherit []string
	Source  string
	Path    string
	Type    string
	Branch  string
	Tag     string
}

//Component represents a graph item of the Nulecule file
type Component struct {
	Name      string
	Source    string
	Params    []Param
	Artifacts map[string][]ArtifactEntry
}
