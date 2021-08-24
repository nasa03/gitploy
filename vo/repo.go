package vo

const (
	RemoteRepoPermRead  RemoteRepoPerm = "read"
	RemoteRepoPermWrite RemoteRepoPerm = "write"
	RemoteRepoPermAdmin RemoteRepoPerm = "admin"
)

type (
	RemoteRepoPerm string

	RemoteRepo struct {
		ID          string         `json:"id"`
		Namespace   string         `json:"namespace"`
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Perm        RemoteRepoPerm `json:"repo_perm"`
	}
)
