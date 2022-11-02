package config

type Yaml struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Labels    Name   `yaml:"labels"`
}

type Name struct {
	Name string `yaml:"name"`
}

type Spec struct {
	Replicas int      `yaml:"replicas"`
	Selector Name     `yaml:"selector"`
	Template Template `yaml:"template"`
}

type Template struct {
	Metadata TempMetadata `yaml:"metadata"`
	Spec     TempSpec     `yaml:"spec"`
}

type TempMetadata struct {
	Labels Name `yaml:"labels"`
}

type TempSpec struct {
	Containers []Containers `yaml:"containers"`
}

type Containers struct {
	Name            string    `yaml:"name"`
	Image           string    `yaml:"image"`
	ImagePullPolicy string    `yaml:"imagePullPolicy"`
	Resources       Resources `yaml:"resources"`
	Ports           []Ports   `yaml:"ports"`
	Env             []Env     `yaml:"env"`
}

type Resources struct {
	Limits   Res `yaml:"limits"`
	Requests Res `yaml:"requests"`
}

type Res struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type Ports struct {
	ContainerPort int `yaml:"containerPort"`
}

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
