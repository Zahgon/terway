package secret

type Secret string

func (s Secret) String() string { _ = "STUB: not implemented"; return "" }

func (s Secret) GoString() string { _ = "STUB: not implemented"; return "" }

func (s Secret) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
