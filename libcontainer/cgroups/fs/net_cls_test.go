// +build linux

package fs

import (
	"strconv"
	"testing"

	"github.com/opencontainers/runc/libcontainer/cgroups/fscommon"
)

const (
	classidBefore = 0x100002
	classidAfter  = 0x100001
)

func TestNetClsSetClassid(t *testing.T) {
	helper := NewCgroupTestUtil("net_cls", t)

	helper.writeFileContents(map[string]string{
		"net_cls.classid": strconv.FormatUint(classidBefore, 10),
	})

	helper.CgroupData.config.Resources.NetClsClassid = classidAfter
	netcls := &NetClsGroup{}
	if err := netcls.Set(helper.CgroupPath, helper.CgroupData.config.Resources); err != nil {
		t.Fatal(err)
	}

	// As we are in mock environment, we can't get correct value of classid from
	// net_cls.classid.
	// So. we just judge if we successfully write classid into file
	value, err := fscommon.GetCgroupParamUint(helper.CgroupPath, "net_cls.classid")
	if err != nil {
		t.Fatal(err)
	}
	if value != classidAfter {
		t.Fatal("Got the wrong value, set net_cls.classid failed.")
	}
}
