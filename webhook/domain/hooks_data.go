package domain

// HooksData defines the data schema retrievable by PipeCD webhooks.
type HooksData struct {
	Type     int      `json:"Type"`
	Metadata Metadata `json:"Metadata"`
}

// Metadata defines the metadata information.
type Metadata struct {
	Deployment Deployment `json:"deployment"`
}

// Deployment defines the deployment information.
type Deployment struct {
	ID               string            `json:"id"`
	ApplicationID    string            `json:"application_id"`
	ApplicationName  string            `json:"application_name"`
	PipedID          string            `json:"piped_id"`
	ProjectID        string            `json:"project_id"`
	GitPath          GitPath           `json:"git_path"`
	CloudProvider    string            `json:"cloud_provider"`
	PlatformProvider string            `json:"platform_provider"`
	Labels           map[string]string `json:"labels"`
	Trigger          Trigger           `json:"trigger"`
	StatusReason     string            `json:"status_reason"`
	CreatedAt        int               `json:"created_at"`
	UpdatedAt        int               `json:"updated_at"`
}

// GitPath defines the GitHub repository path information.
type GitPath struct {
	Repo           Repo   `json:"repo"`
	Path           string `json:"path"`
	ConfigFilename string `json:"config_filename"`
	URL            string `json:"url"`
}

// Repo defines the repository information.
type Repo struct {
	ID     string `json:"id"`
	Remote string `json:"remote"`
	Branch string `json:"branch"`
}

// Trigger defines the trigger information.
type Trigger struct {
	Commit    Commit `json:"commit"`
	Timestamp int    `json:"timestamp"`
}

// Commit defines the commit information.
type Commit struct {
	Hash      string `json:"hash"`
	Message   string `json:"message"`
	Author    string `json:"author"`
	Branch    string `json:"branch"`
	URL       string `json:"url"`
	CreatedAt int    `json:"created_at"`
}
