package releases

type Build struct {
	Arch        string `json:"arch"`
	Os          string `json:"os"`
	Unsupported bool   `json:"unsupported"`
	Url         string `json:"url"`
}
