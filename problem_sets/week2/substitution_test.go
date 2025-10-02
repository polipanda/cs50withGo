package week2

import "testing"

func TestSubstitutionCipherText(t *testing.T) {
	for _, scenario := range substitutionScenarios() {
		result, err := substituteCipherText(scenario.plaintext, scenario.key)
		if err != nil && !scenario.expectError {
			t.Errorf("should not have gotten err %v", err)
		}
		if result != scenario.ciphertext {
			t.Errorf("expected %s, got %s", scenario.ciphertext, result)
		}
	}
}

type substitutionScenario struct {
	plaintext   string
	key         string
	ciphertext  string
	expectError bool
}

func substitutionScenarios() []substitutionScenario {
	return []substitutionScenario{
		{
			key:         "werrt",
			expectError: true,
		},
		{
			key:         "abcdefghijklmnopqrstuvwxyz!@#",
			expectError: true,
		},
		{
			key:         "abcdefghijklmnopqrstuvwx!@",
			expectError: true,
		},
		{
			key:         "abcdefghijklmnopqrstuvwxzz",
			expectError: true,
		},
		{
			plaintext:   "Hello",
			key:         "NQXPOMAFTRHLZGECYJIUWSKDVB",
			ciphertext:  "Folle",
			expectError: false,
		},
		{
			plaintext: `In Solitude, you may hear a rumor from innkeepers that a mad beggar is wandering around the
				streets of the city, distressed over the absence of his master, saying things like: "Please, take pity 
				on an old madman!" and "You! You'll help me! You help people, right? That's what you do?"`,
			key: "zyqrbicnshavxptmfkjldguewo",
			ciphertext: `Sp Jtvsldrb, wtd xzw nbzk z kdxtk iktx sppabbmbkj lnzl z xzr ybcczk sj uzprbkspc zktdpr lnb
				jlkbblj ti lnb qslw, rsjlkbjjbr tgbk lnb zyjbpqb ti nsj xzjlbk, jzwspc lnspcj vsab: "Mvbzjb, lzab mslw 
				tp zp tvr xzrxzp!" zpr "Wtd! Wtd'vv nbvm xb! Wtd nbvm mbtmvb, kscnl? Lnzl'j unzl wtd rt?"`,

			expectError: false,
		},
	}
}
