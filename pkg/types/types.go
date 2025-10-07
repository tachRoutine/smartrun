package types

type LLMResponse struct {
    Instructions string    `json:"instructions"`
    Commands     []Command `json:"commands"`
}

type Command struct {
    Command     string `json:"command"`
    Description string `json:"description"`
    Platform    string `json:"platform"` // "windows", "linux", "macos", "all"
    Dangerous   bool   `json:"dangerous"`
}
