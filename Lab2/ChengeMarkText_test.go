package lab2

import (
	"fmt"
	"testing"
)

func TestChengeText(t *testing.T) {
	for _, tc := range []struct {
		nameTest string
		textIn   string
		resText  string
	}{
		{nameTest: "1", textIn: "Test", resText: "<p>Test</p>"},
		{nameTest: "2", textIn: "**Test**", resText: "<p><b>Test</b></p>"},
		{nameTest: "3", textIn: "_Test_", resText: "<p><i>Test</i></p>"},
		{nameTest: "4", textIn: "`Test`", resText: "<p><tt>Test</tt></p>"},
		{nameTest: "5", textIn: "```Test```", resText: "<p><pre>Test</pre></p>"},
		{nameTest: "6", textIn: "T  e  s  t", resText: "<p>T</p><p>e</p><p>s</p><p>t</p>"},
		{nameTest: "7", textIn: "```T*e`s_t```", resText: "<p><pre>T*e`s_t</pre></p>"},
		{nameTest: "8", textIn: "_ Test", resText: "<p>_ Test</p>"},
		{nameTest: "9", textIn: "Te_st", resText: "<p>Te_st</p>"},
		{nameTest: "10", textIn: "**Test*", resText: "<p><b>Test"},  // тест падає через помилку у тексті
		{nameTest: "11", textIn: "_Test  _", resText: "<p><i>Test"}, // тест падає через помилку у тексті
	} {
		t.Run(tc.nameTest, func(t *testing.T) {
			if got, textErr := ChengeText(tc.textIn); textErr != "" {
				t.Errorf("Have error: %s. ResText: %s, got: %s", textErr, tc.resText, got)
			} else {
				fmt.Println("ResText:", tc.resText, "got:", got)
			}
		})
	}
}
