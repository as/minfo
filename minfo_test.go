package minfo

import (
	"encoding/hex"
	"strings"
)

/*
func TestReadStdin(t *testing.T) {
	file, err := Read(bytes.NewReader(h264AnnexBBitstream))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(file)
}
*/

// h264AnnexBBitstream is a compliant h264 bitstream. This stream can be
// decoded in ffmpeg and outputs the Lenna test image.
var h264AnnexBBitstream, _ = hex.DecodeString(clean(`
000000016764000aac7284442684000003000400000300ca3c48961180
0000000168e8438f13213
000000165
888100054e7f87df61a58b95eea4e938b76a306a71b955600b762eb50ee48059
27b867a963375e822055fbe46ae9373572e222919e4dff6086ce7e42b795ce2a
e126be87738426ba1636f4e69f17dad8647554b1f3450c0b3c74b39dbceb5373
87c30e62474862ca59eb863f3afa86b5bfa86d06165082c4ce629e4ee64cc730
3edea10bd8830bb6b828bca9eb7743fc7a17948521ca376b3095b546773060b7
12d68cc5548529d869a96f124e71dfe3e2b16b6bbf9ffb2e5730a96976c446a2
dffa91d95074551d49045a1cd686687cb661486c96e6124c27adbac751998ed0
f0ed8ef6657979a612a195dbc8aee3b635e68dbc48a37faf4a288a53e27e6808
9f67779852db5084d65e25e14a995834c711d643ffc4fd9a4416d1b2fb02dba1
896934c2325598f99bb2313f49590c068cdba5b29d7e122fd0879444e40a76ef
992d911839503b293bf52c9773489183b0a6f34b702f1c8f3b7823c6aa864643
1dd72a235e2cd9480af5f52cd1fb3ff04b7837e945dd72cf8035c39507f3d906
e54a5876036c81206245654473bcfec19f31e5db895c6b79d86890d726a8a188
8681dc9a4f40a523c7debe6f76ab7916512167832ef3d6271a42c294d15d6cdb
4a7ae2cb0bb0680bbe19590050fcc0bd9df5f5f8a81719d6b3e974ba50e52c45
7bf993ea5af9a930b16f5b36241e8d5557f4cc67b2656aa93626d006b8e2e373
8bd1c01c5215cab5ac603e3642f12cbd9977aba8a9a48e9c8b84de73f0912997
aedbafd6f85e9b86b3b303b3ac756fa611692f3d3acefa538660956cbbc54ef3
`))

func clean(s string) string {
	for _, rm := range []string{"\n", "\t", " "} {
		s = strings.Replace(s, rm, "", -1)
	}
	return s
}