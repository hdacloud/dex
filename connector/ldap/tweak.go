package ldap

import (
	"strings"

	"github.com/dexidp/dex/connector"
)

// TweakIdentity adjusts attributes received from the LDAP directory. Don't ask
// why this is necessary. Just learn to accept it.
func tweakIdentity(id connector.Identity) connector.Identity {
	id.Username = tweakName(id)
	return id
}

func tweakName(id connector.Identity) string {
	name := id.Username
	if name == " " {
		return id.PreferredUsername
	}

	xs := strings.Split(name, ", ")
	if len(xs) == 1 {
		return name
	}

	if strings.Contains(xs[1], " (") {
		xs[1] = strings.Split(xs[1], " (")[0]
	}

	return xs[1] + " " + xs[0]
}
