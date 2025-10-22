package output

// importsMap maps template names to their required import paths.
// We do not need support for aliases in import for now.
var importsMap = map[string]string{
	"testify": "github.com/stretchr/testify/assert",
}
