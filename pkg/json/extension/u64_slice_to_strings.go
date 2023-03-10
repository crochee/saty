package extension

import (
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/json-iterator/go"
)

type u64SliceAsStringsCodec struct {
	jsoniter.DummyExtension
}

func (extension *u64SliceAsStringsCodec) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, binding := range structDescriptor.Fields {
		if binding.Field.Type().Kind() == reflect.Slice {
			tagParts := strings.Split(binding.Field.Tag().Get("json"), ",")
			if len(tagParts) <= 1 {
				continue
			}
			for _, tagPart := range tagParts[1:] {
				if tagPart == "strings" {
					binding.Encoder = &funcEncoder{fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {
						nums := *((*[]uint64)(ptr))
						strs := make([]string, 0, len(nums))
						for _, num := range nums {
							strs = append(strs, strconv.FormatUint(num, 10))
						}
						stream.WriteVal(strs)
					}}
					binding.Decoder = &funcDecoder{}
					break
				}
			}
		}
	}
}