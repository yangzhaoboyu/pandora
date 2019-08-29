package common

import "github.com/koding/multiconfig"

//Load configuration file
func LoadConfig(path string, model interface{}) error {
	loader := multiconfig.NewWithPath(path)
	err := loader.Load(model)
	if err != nil {
		return err
	}
	return nil
}
