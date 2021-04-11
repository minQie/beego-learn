package child

type ResponseCacheConfig struct {
	MaxEntries        int   `yaml:"max_entries"`
	ExpireTimeSeconds int64 `yaml:"expire_time_seconds"`
}
