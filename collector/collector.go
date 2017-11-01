
package collector

import (
	"strings"
)

func normalizeFQN(fqn string) string {
	var _fqn string = fqn
	_fqn = strings.Replace(_fqn, ".", "_", -1)
	_fqn = strings.Replace(_fqn, "-", "_", -1)

	return _fqn
}
