// Code generated by "stringer -type=MMBearerIpFamily"; DO NOT EDIT.

package enums

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmBearerIpFamilyNone-0]
	_ = x[MmBearerIpFamilyIpv4-1]
	_ = x[MmBearerIpFamilyIpv6-2]
	_ = x[MmBearerIpFamilyIpv4v6-4]
	_ = x[MmBearerIpFamilyAny-4294967295]
}

const (
	_MMBearerIpFamily_name_0 = "MmBearerIpFamilyNoneMmBearerIpFamilyIpv4MmBearerIpFamilyIpv6"
	_MMBearerIpFamily_name_1 = "MmBearerIpFamilyIpv4v6"
	_MMBearerIpFamily_name_2 = "MmBearerIpFamilyAny"
)

var (
	_MMBearerIpFamily_index_0 = [...]uint8{0, 20, 40, 60}
)

func (i MMBearerIpFamily) String() string {
	switch {
	case i <= 2:
		return _MMBearerIpFamily_name_0[_MMBearerIpFamily_index_0[i]:_MMBearerIpFamily_index_0[i+1]]
	case i == 4:
		return _MMBearerIpFamily_name_1
	case i == 4294967295:
		return _MMBearerIpFamily_name_2
	default:
		return "MMBearerIpFamily(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
