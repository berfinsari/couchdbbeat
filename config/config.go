// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

type CouchdbbeatConfig struct {
	Period *int64
	Port   *string
	Host   *string
}

type ConfigSettings struct {
	Input CouchdbbeatConfig
}
