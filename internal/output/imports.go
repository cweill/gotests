package output

const GoMonkeyPkg = "gomonkey"

// we do not need support for aliases in import for now.
var importsMap = map[string]string{
	"testify": "github.com/stretchr/testify/assert",
	GoMonkeyPkg: "github.com/agiledragon/gomonkey/v2",
}
