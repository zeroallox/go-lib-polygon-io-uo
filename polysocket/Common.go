package polysocket

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.Config{
	EscapeHTML:                    false,
	ObjectFieldMustBeSimpleString: true,
}.Froze()
