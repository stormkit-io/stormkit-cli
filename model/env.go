package model

// Env model rappresent an environemt of an application
type Env struct {
	ID     string `json:"id"`
	AppID  string `json:"appId"`
	Env    string `json:"env"`
	Branch string `json:"master"`
	Build  struct {
		Entry string `json:"entry"`
		Cmd   string `json:"cmd"`
	}
	AutoPublish bool `json:"autoPublish"`
	LastDeploy  struct {
		ID        int   `json:"id"`
		CreatedAt int64 `json:"createdAt"`
		Exit      int   `json:"exit"`
	}
	Domain struct {
		Verified bool `json:"verified"`
	}
}

// EnvsArray array of environments of an application
type EnvsArray struct {
	Envs []Env `json:"envs"`
}

// Names returns the names of the envs
func (a *EnvsArray) Names() []string {
	n := []string{}

	for _, v := range a.Envs {
		n = append(n, v.Env)
	}

	return n
}

// MockEnv for run tests
var MockEnv = Env{
	ID:          "1234",
	AppID:       "1235",
	Env:         "env",
	Branch:      "branch",
	AutoPublish: true,
}

// MockEnvsArray for run tests
var MockEnvsArray = EnvsArray{
	Envs: []Env{
		MockEnv,
		Env{
			ID:          "1236",
			AppID:       "1235",
			Env:         "env2",
			Branch:      "branch2",
			AutoPublish: false,
		},
	},
}
