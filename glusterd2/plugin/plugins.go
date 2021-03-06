// +build plugins

package plugin

import (
	"github.com/gluster/glusterd2/plugins/bitrot"
	"github.com/gluster/glusterd2/plugins/georeplication"
	"github.com/gluster/glusterd2/plugins/hello"
	"github.com/gluster/glusterd2/plugins/quota"
)

// PluginsList is a list of plugins which implements GlusterdPlugin interface
var PluginsList = []GlusterdPlugin{
	&hello.Plugin{},
	&georeplication.Plugin{},
	&bitrot.Plugin{},
	&quota.Plugin{},
}
