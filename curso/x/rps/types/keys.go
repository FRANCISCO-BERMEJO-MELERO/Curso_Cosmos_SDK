package types

const (
	// ModuleName defines the module name
	ModuleName = "rps"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rps"
)

var (
	ParamsKey = []byte("p_rps")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
