// +build !linux

package daemon // import "github.com/sdslabs/docker/daemon"

func selinuxSetDisabled() {
}

func selinuxFreeLxcContexts(label string) {
}

func selinuxEnabled() bool {
	return false
}
