package git

import "github.com/pkg/errors"

// Save saves a file to a branch.
func Save(branch, user, path string) error {
	list := makeConfigCmd(user)
	list = append(list, makeSaveCmd(branch, path)...)
	for _, cmd := range list {
		if err := runCmd(cmd); err != nil {
			return errors.Wrapf(err, "error running command: %s", cmd)
		}
	}
	return nil
}
