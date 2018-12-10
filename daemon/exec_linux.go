package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"github.com/sdslabs/docker/container"
	"github.com/sdslabs/docker/daemon/caps"
	"github.com/sdslabs/docker/daemon/exec"
	"github.com/opencontainers/runc/libcontainer/apparmor"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func (daemon *Daemon) execSetPlatformOpt(c *container.Container, ec *exec.Config, p *specs.Process) error {
	if len(ec.User) > 0 {
		uid, gid, additionalGids, err := getUser(c, ec.User)
		if err != nil {
			return err
		}
		p.User = specs.User{
			UID:            uid,
			GID:            gid,
			AdditionalGids: additionalGids,
		}
	}
	if ec.Privileged {
		if p.Capabilities == nil {
			p.Capabilities = &specs.LinuxCapabilities{}
		}
		p.Capabilities.Bounding = caps.GetAllCapabilities()
		p.Capabilities.Permitted = p.Capabilities.Bounding
		p.Capabilities.Inheritable = p.Capabilities.Bounding
		p.Capabilities.Effective = p.Capabilities.Bounding
	}
	if apparmor.IsEnabled() {
		var appArmorProfile string
		if c.AppArmorProfile != "" {
			appArmorProfile = c.AppArmorProfile
		} else if c.HostConfig.Privileged {
			// `docker exec --privileged` does not currently disable AppArmor
			// profiles. Privileged configuration of the container is inherited
			appArmorProfile = "unconfined"
		} else {
			appArmorProfile = "docker-default"
		}

		if appArmorProfile == "docker-default" {
			// Unattended upgrades and other fun services can unload AppArmor
			// profiles inadvertently. Since we cannot store our profile in
			// /etc/apparmor.d, nor can we practically add other ways of
			// telling the system to keep our profile loaded, in order to make
			// sure that we keep the default profile enabled we dynamically
			// reload it if necessary.
			if err := ensureDefaultAppArmorProfile(); err != nil {
				return err
			}
		}
		p.ApparmorProfile = appArmorProfile
	}
	daemon.setRlimits(&specs.Spec{Process: p}, c)
	return nil
}
