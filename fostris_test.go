package fostris_test

import (
	"testing"

	"github.com/phille97/fostris"
)

func TestEncodeAndDecodeMatch(t *testing.T) {
	alts := []string{
		"I like creamy peanut butter",
		"Foo Bar",
		"F O S T R I S",
		"FOSTRIS",
		"SIRTSOF",
	}
	for _, alt := range alts {
		altEnc := fostris.Encode(alt)
		altDec := fostris.Decode(altEnc)
		if altDec != alt {
			t.Errorf("encode/decode failed to match \"%s\" with \"%s\"", alt, altDec)
		}
	}
}

func TestEncodeAndDecodeStaticExamples(t *testing.T) {
	encoded := "fosTRisfoStrIS fOSTriSfoStRIsfoSTriSfoSTRIs foStriSfosTRiS fOsTrisfostRiSfosTRiSfOsTRiSfOStRiSfoStriSfosTRiS"
	decoded := "My name is Fostris"

	if fostris.Encode(decoded) != encoded {
		t.Errorf("encode failed with our static example")
	}

	if fostris.Decode(encoded) != decoded {
		t.Errorf("decode failed with our static example")
	}
}
