package templates

import "embed"

//go:embed test/* testify/*
var FS embed.FS
