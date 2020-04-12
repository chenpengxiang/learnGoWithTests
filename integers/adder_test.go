package integers

import "testing"


func assertCorrectMessage(t *testing.T, expected, result int) {
	if expected != result {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}

func TestAdder(t *testing.T) {

	t.Run("2 + 2 = 4", func (t *testing.T){
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, expected, sum)
	})

	t.Run("5 + 2 = 7", func (t *testing.T)  {
		sum := Add(5, 2)
		expected := 7
		assertCorrectMessage(t, expected, sum)
	})
}