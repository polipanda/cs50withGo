package week2

import "testing"

func TestCipherText(t *testing.T) {
	for _, scenario := range caesarScenarios() {
		result, err := cipherText(scenario.plaintext, scenario.key)
		if err != nil && !scenario.expectError {
			t.Errorf("should not have gotten err %v", err)
		}
		if result != scenario.ciphertext {
			t.Errorf("expected %s, got %s", scenario.ciphertext, result)
		}
	}
}

type caesarScenario struct {
	plaintext   string
	key         int
	ciphertext  string
	expectError bool
}

func caesarScenarios() []caesarScenario {
	return []caesarScenario{
		{
			plaintext:   "hello",
			key:         -10,
			ciphertext:  "",
			expectError: true,
		},
		{
			plaintext:   "hello!",
			key:         10,
			ciphertext:  "rovvy!",
			expectError: false,
		},
		{
			plaintext:   "let's go to the moon!",
			key:         999,
			ciphertext:  "wpe'd rz ez esp xzzy!",
			expectError: false,
		},
		{
			plaintext:   "SuPPOse JUMbles are TRENDY atm",
			key:         150,
			ciphertext:  "MoJJImy DOGvfym uly NLYHXS ung",
			expectError: false,
		},
	}
}
