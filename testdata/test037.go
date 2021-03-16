package testdata

import "net/url"

type someIndirectImportedStruct url.URL

func (smtg *someIndirectImportedStruct) Foo037() {}
