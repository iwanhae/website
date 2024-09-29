// Code generated by "stringer -type=DocumentType -output=post_types.go -linecomment"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DocumentTypeUnknown-0]
	_ = x[DocumentTypeMarkdown-1]
	_ = x[DocumentTypeHTML-2]
}

const _DocumentType_name = "unknownmarkdownhtml"

var _DocumentType_index = [...]uint8{0, 7, 15, 19}

func (i DocumentType) String() string {
	if i >= DocumentType(len(_DocumentType_index)-1) {
		return "DocumentType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DocumentType_name[_DocumentType_index[i]:_DocumentType_index[i+1]]
}