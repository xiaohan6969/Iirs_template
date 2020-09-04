package corsServer

import (
	asconfig "../../config"
	"../../config/configHelper"
	"github.com/rs/cors"
)

func Cors() *cors.Cors {
	if asconfig.Config.Has("cors") {
		config := configHelper.GetTree(asconfig.Config, "cors")
		options := cors.Options{
			AllowedOrigins:     configHelper.GetStringArray(config, "allowed-origins"),
			AllowedMethods:     configHelper.GetStringArray(config, "allowed-methods"),
			AllowedHeaders:     configHelper.GetStringArray(config, "allowed-headers"),
			ExposedHeaders:     configHelper.GetStringArray(config, "exposed-methods"),
			AllowCredentials:   configHelper.GetBool(config, "allow-credentials"),
			OptionsPassthrough: configHelper.GetBool(config, "options-passthrough"),
			Debug:              configHelper.GetBool(config, "debug", false),
		}
		return cors.New(options)
	}
	return cors.Default()
}
