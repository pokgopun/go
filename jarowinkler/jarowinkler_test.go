package jarowinkler

import "testing"

func TestSimilarity(t *testing.T) {
	data := []struct {
		str1, str2 string
		sim0, sim  float64
	}{
		{"NICHLESON", "NICHULSON", 0.926, 0.956},
		{"JONES", "JOHNSON", 0.790, 0.832},
		{"MASSEY", "MASSIE", 0.889, 0.933},
		{"ABROMS", "ABRAMS", 0.889, 0.922},
		{"JERALDINE", "GERALDINE", 0.926, 0.926},
		{"MARHTA", "MARTHA", 0.944, 0.961},
		{"MICHELLE", "MICHAEL", 0.869, 0.921},
		{"JULIES", "JULIUS", 0.889, 0.933},
		{"TANYA", "TONYA", 0.867, 0.880},
		{"DWAYNE", "DUANE", 0.822, 0.840},
		{"SEAN", "SUSAN", 0.783, 0.805},
		{"JON", "JOHN", 0.917, 0.933},
	}
	for _, d := range data {
		jw := New(d.str1, d.str2)
		if res := jw.Similarity(false); res != d.sim0 {
			t.Errorf("incorrect result for %q %q, expected %.3f, got %.3f", d.str1, d.str2, d.sim0, res)
		}
		if res := jw.Similarity(true); res != d.sim {
			t.Errorf("incorrect result for %q %q, expected %.3f, got %.3f", d.str1, d.str2, d.sim, res)
		}
	}
}
