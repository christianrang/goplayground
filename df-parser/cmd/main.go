package main

import (
	dfparser "dfparser/pkg/parser"
	"fmt"
)

const (
	dfoutput string = `
Filesystem          Size  Used Avail Use% Mounted on
dev                  16G     0   16G   0% /dev
run                  16G  2.4M   16G   1% /run
/dev/sda1           196G   19G  167G  11% /
tmpfs                16G   39M   16G   1% /dev/shm
tmpfs                16G   14M   16G   1% /tmp
/dev/loop6          128K  128K     0 100% /var/lib/snapd/snap/bare/5
/dev/loop3          252M  252M     0 100% /var/lib/snapd/snap/brave/153
/dev/loop1          111M  111M     0 100% /var/lib/snapd/snap/core/12834
/dev/loop4           66M   66M     0 100% /var/lib/snapd/snap/gtk-common-themes/1519
/dev/loop5           56M   56M     0 100% /var/lib/snapd/snap/core18/2344
/dev/loop8           37M   37M     0 100% /var/lib/snapd/snap/gh/502
/dev/loop0           44M   44M     0 100% /var/lib/snapd/snap/snapd/15177
/dev/loop7          252M  252M     0 100% /var/lib/snapd/snap/brave/157
`
)

func main() {
	disks := dfparser.Parse(dfoutput)
	for _, disk := range disks {
		fmt.Println(disk)
	}
}
