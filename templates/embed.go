package templates

import "embed"

//go:embed test/* test_empty/* testify/*
var FS embed.FS
