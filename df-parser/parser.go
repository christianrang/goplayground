package dfparser

type Disk struct {
	Name          string
	Size          string
	Available     string
	UsePercentage string
	MountPoint    string
}

func AppendNewDisk(disks []Disk, newDisk Disk) []Disk {
	if len(newDisk.Name) <= 0 {
		return disks
	}

	if newDisk.Name == "Filesystem" {
		return disks
	}

	return append(disks, newDisk)
}

func Parse(output string) []Disk {
	var (
		disks []Disk
		tmp   = Disk{}
		// Tracks the column we are currently parsing from the df table
		valueIdx = 0
	)

	for i := 0; i < len(output); i++ {
		switch char := output[i]; char {
		case '\n':
			disks = AppendNewDisk(disks, tmp)
			tmp = Disk{}
			valueIdx = 0
		case ' ':
			for i < len(output)-1 && output[i+1] == ' ' {
				i++
			}
		default:
			word := []byte{}
			for i < len(output) && output[i] != ' ' && output[i] != '\n' {
				word = append(word, output[i])
				if i < len(output) && output[i+1] != ' ' && output[i+1] != '\n' {
					i++
				} else {
					break
				}
			}
			tmp.setDiskValue(valueIdx, string(word))
			valueIdx++
		}
	}

	return disks
}

func (disk *Disk) setDiskValue(idx int, value string) {
	switch idx {
	case 0:
		disk.Name = value
	case 1:
		disk.Size = value
	case 2:
		disk.Available = value
	case 3:
		disk.UsePercentage = value
	case 4:
		disk.MountPoint = value
	}
}
