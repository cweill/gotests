package templates

import "embed"

// FS contains embedded template files for test generation.
//
//go:embed test/* testify/*
var FS embed.FS
