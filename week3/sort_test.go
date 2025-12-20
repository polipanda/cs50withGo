package week3

import (
	"os/exec"
	"testing"
)

func TestSortSelections(t *testing.T) {
	t.Run("check sort1", func(t *testing.T) {
		if result := returnAnswer(SORT1); result != BUBBLE {
			t.Errorf("expected BUBBLE answer, got %s", result)
		}
	})

	t.Run("check sort2", func(t *testing.T) {
		if result := returnAnswer(SORT2); result != MERGE {
			t.Errorf("expected BUBBLE answer, got %s", result)
		}
	})

	t.Run("check sort3", func(t *testing.T) {
		if result := returnAnswer(SORT3); result != SELECTION {
			t.Errorf("expected BUBBLE answer, got %s", result)
		}
	})
}

func BenchmarkBinaries(b *testing.B) {
	sort1path := "./sort_problem_files/sort1"
	sort2path := "./sort_problem_files/sort2"
	sort3path := "./sort_problem_files/sort3"
	tests := []struct {
		name     string
		location string
		args     []string
	}{
		{"sort1_random5000", sort1path, []string{"./sort_problem_files/random5000.txt"}},
		{"sort1_random10000", sort1path, []string{"./sort_problem_files/random10000.txt"}},
		{"sort1_random50000", sort1path, []string{"./sort_problem_files/random50000.txt"}},
		{"sort1_reversed5000", sort1path, []string{"./sort_problem_files/reversed5000.txt"}},
		{"sort1_reversed10000", sort1path, []string{"./sort_problem_files/reversed10000.txt"}},
		{"sort1_reversed50000", sort1path, []string{"./sort_problem_files/reversed50000.txt"}},
		{"sort1_sorted5000", sort1path, []string{"./sort_problem_files/sorted5000.txt"}},
		{"sort1_sorted10000", sort1path, []string{"./sort_problem_files/sorted10000.txt"}},
		{"sort1_sorted50000", sort1path, []string{"./sort_problem_files/sorted50000.txt"}},

		{"sort2_random5000", sort2path, []string{"./sort_problem_files/random5000.txt"}},
		{"sort2_random10000", sort2path, []string{"./sort_problem_files/random10000.txt"}},
		{"sort2_random50000", sort2path, []string{"./sort_problem_files/random50000.txt"}},
		{"sort2_reversed5000", sort2path, []string{"./sort_problem_files/reversed5000.txt"}},
		{"sort2_reversed10000", sort2path, []string{"./sort_problem_files/reversed10000.txt"}},
		{"sort2_reversed50000", sort2path, []string{"./sort_problem_files/reversed50000.txt"}},
		{"sort2_sorted5000", sort2path, []string{"./sort_problem_files/sorted5000.txt"}},
		{"sort2_sorted10000", sort2path, []string{"./sort_problem_files/sorted10000.txt"}},
		{"sort2_sorted50000", sort2path, []string{"./sort_problem_files/sorted50000.txt"}},

		{"sort3_random5000", sort3path, []string{"./sort_problem_files/random5000.txt"}},
		{"sort3_random10000", sort3path, []string{"./sort_problem_files/random10000.txt"}},
		{"sort3_random50000", sort3path, []string{"./sort_problem_files/random50000.txt"}},
		{"sort3_reversed5000", sort3path, []string{"./sort_problem_files/reversed5000.txt"}},
		{"sort3_reversed10000", sort3path, []string{"./sort_problem_files/reversed10000.txt"}},
		{"sort3_reversed50000", sort3path, []string{"./sort_problem_files/reversed50000.txt"}},
		{"sort3_sorted5000", sort3path, []string{"./sort_problem_files/sorted5000.txt"}},
		{"sort3_sorted10000", sort3path, []string{"./sort_problem_files/sorted10000.txt"}},
		{"sort3_sorted50000", sort3path, []string{"./sort_problem_files/sorted50000.txt"}},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				cmd := exec.Command(test.location, test.args...)
				if err := cmd.Run(); err != nil {
					b.Fatalf("%s failed: %v", test.name, err)
				}
			}
		})
	}
}
