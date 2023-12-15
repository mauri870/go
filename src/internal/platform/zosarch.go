// Code generated by go test internal/platform -fix. DO NOT EDIT.

// To change the information in this file, edit the cgoEnabled and/or firstClass
// maps in cmd/dist/build.go, then run 'go generate internal/platform'.

package platform

// List is the list of all valid GOOS/GOARCH combinations,
// including known-broken ports.
var List = []OSArch{
	{"aix", "ppc64"},
	{"android", "386"},
	{"android", "amd64"},
	{"android", "arm"},
	{"android", "arm64"},
	{"darwin", "amd64"},
	{"darwin", "arm64"},
	{"dragonfly", "amd64"},
	{"freebsd", "386"},
	{"freebsd", "amd64"},
	{"freebsd", "arm"},
	{"freebsd", "arm64"},
	{"freebsd", "riscv64"},
	{"illumos", "amd64"},
	{"ios", "amd64"},
	{"ios", "arm64"},
	{"js", "wasm"},
	{"linux", "386"},
	{"linux", "amd64"},
	{"linux", "arm"},
	{"linux", "arm64"},
	{"linux", "loong64"},
	{"linux", "mips"},
	{"linux", "mips64"},
	{"linux", "mips64le"},
	{"linux", "mipsle"},
	{"linux", "ppc64"},
	{"linux", "ppc64le"},
	{"linux", "riscv64"},
	{"linux", "s390x"},
	{"linux", "sparc64"},
	{"netbsd", "386"},
	{"netbsd", "amd64"},
	{"netbsd", "arm"},
	{"netbsd", "arm64"},
	{"openbsd", "386"},
	{"openbsd", "amd64"},
	{"openbsd", "arm"},
	{"openbsd", "arm64"},
	{"openbsd", "mips64"},
	{"openbsd", "ppc64"},
	{"openbsd", "riscv64"},
	{"plan9", "386"},
	{"plan9", "amd64"},
	{"plan9", "arm"},
	{"solaris", "amd64"},
	{"wasip1", "wasm"},
	{"wasip1", "wasm32"},
	{"windows", "386"},
	{"windows", "amd64"},
	{"windows", "arm"},
	{"windows", "arm64"},
}

var distInfo = map[OSArch]osArchInfo{
	{"aix", "ppc64"}:       {CgoSupported: true},
	{"android", "386"}:     {CgoSupported: true},
	{"android", "amd64"}:   {CgoSupported: true},
	{"android", "arm"}:     {CgoSupported: true},
	{"android", "arm64"}:   {CgoSupported: true},
	{"darwin", "amd64"}:    {CgoSupported: true, FirstClass: true},
	{"darwin", "arm64"}:    {CgoSupported: true, FirstClass: true},
	{"dragonfly", "amd64"}: {CgoSupported: true},
	{"freebsd", "386"}:     {CgoSupported: true},
	{"freebsd", "amd64"}:   {CgoSupported: true},
	{"freebsd", "arm"}:     {CgoSupported: true},
	{"freebsd", "arm64"}:   {CgoSupported: true},
	{"freebsd", "riscv64"}: {CgoSupported: true},
	{"illumos", "amd64"}:   {CgoSupported: true},
	{"ios", "amd64"}:       {CgoSupported: true},
	{"ios", "arm64"}:       {CgoSupported: true},
	{"js", "wasm"}:         {},
	{"linux", "386"}:       {CgoSupported: true, FirstClass: true},
	{"linux", "amd64"}:     {CgoSupported: true, FirstClass: true},
	{"linux", "arm"}:       {CgoSupported: true, FirstClass: true},
	{"linux", "arm64"}:     {CgoSupported: true, FirstClass: true},
	{"linux", "loong64"}:   {CgoSupported: true},
	{"linux", "mips"}:      {CgoSupported: true},
	{"linux", "mips64"}:    {CgoSupported: true},
	{"linux", "mips64le"}:  {CgoSupported: true},
	{"linux", "mipsle"}:    {CgoSupported: true},
	{"linux", "ppc64"}:     {},
	{"linux", "ppc64le"}:   {CgoSupported: true},
	{"linux", "riscv64"}:   {CgoSupported: true},
	{"linux", "s390x"}:     {CgoSupported: true},
	{"linux", "sparc64"}:   {CgoSupported: true, Broken: true},
	{"netbsd", "386"}:      {CgoSupported: true},
	{"netbsd", "amd64"}:    {CgoSupported: true},
	{"netbsd", "arm"}:      {CgoSupported: true},
	{"netbsd", "arm64"}:    {CgoSupported: true},
	{"openbsd", "386"}:     {CgoSupported: true},
	{"openbsd", "amd64"}:   {CgoSupported: true},
	{"openbsd", "arm"}:     {CgoSupported: true},
	{"openbsd", "arm64"}:   {CgoSupported: true},
	{"openbsd", "mips64"}:  {CgoSupported: true, Broken: true},
	{"openbsd", "ppc64"}:   {},
	{"openbsd", "riscv64"}: {Broken: true},
	{"plan9", "386"}:       {},
	{"plan9", "amd64"}:     {},
	{"plan9", "arm"}:       {},
	{"solaris", "amd64"}:   {CgoSupported: true},
	{"wasip1", "wasm"}:     {},
	{"wasip1", "wasm32"}:   {},
	{"windows", "386"}:     {CgoSupported: true, FirstClass: true},
	{"windows", "amd64"}:   {CgoSupported: true, FirstClass: true},
	{"windows", "arm"}:     {},
	{"windows", "arm64"}:   {CgoSupported: true},
}
