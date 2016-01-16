package ip

type IPKind int

const (
	Open IPKind = iota
	Close
	EOF
	Normal
)

type IP struct {
	Kind IPKind
	Data string
}
