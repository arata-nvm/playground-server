package model

type Job struct {
	Stdin   string `json:"input"`
	Program string `json:"program"`
}

type Result struct {
	Error         string `json:"error"`
	Stdout        string `json:"output"`
	CompilerError string `json:"compiler_error"` // Stderr
}
