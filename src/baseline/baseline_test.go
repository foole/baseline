package baseline

import "testing"

type stringPair struct {
    value       string
    expected    string
}

var StringPairs = []stringPair {
    { "1", "1" },
    { "short string", "short string" },
    { "50 character string...............................",
      "50 character string..............................." },
    { "-51 character string...............................",
      "51 character string..............................." },
    { "------------------------------80 character string...............................",
      "80 character string..............................." },
}

func TestTruncation(t *testing.T) {
    for _, pair := range StringPairs {
        trunc := TruncateTestName(pair.value)
        if trunc != pair.expected {
            t.Error(
                "For", pair.value,
                "Expected", pair.expected,
                "Got", trunc,
            )
        }
    }
}
