package pack

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestMain(m *testing.M) {
	fmt.Println("Tests are about to run.")
	result := m.Run()
	fmt.Println("Test suite finishs")
	os.Exit(result)
}

func TestCanAddNumbers(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("Skipping long tests")
	}
	time.Sleep(1 * time.Second)

	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Fatal("Failed to add numbers.")
			//the same as calling
			//t.Log("Failed to add numbers.")
			//t.FailNow()

			//t.Error("OMG")
			//the same as calling
			//t.Log("Failed to add numbers.")
			//t.Fail()
		}
	}
}

func TestCanSubtractNumbers(t *testing.T) {
	t.Parallel()
	result := Subtract(10, 5)
	time.Sleep(1 * time.Second)
	if result != 5 {
		t.Error("Failed to subtract two numbers properly.")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	if testing.Verbose() {
		t.Skip("Not implmented yet.")
	}
}
