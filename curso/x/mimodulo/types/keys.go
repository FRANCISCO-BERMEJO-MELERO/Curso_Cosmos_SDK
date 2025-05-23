package types

const (
	// ModuleName defines the module name
	ModuleName = "mimodulo"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mimodulo"
)

var (
	ParamsKey = []byte("p_mimodulo")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
