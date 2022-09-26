package cryptoformulas

import "testing"

func Test(t *testing.T) {
	got := ProfitPercentageExpectation(1.00, 2.00)
	want := 100.00

	if want != got {
		t.Errorf("want: %.2f, got: %.2f", want, got)
	}
}
