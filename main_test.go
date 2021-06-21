package main

import "testing"

func Test_fooA(t *testing.T) {
	arr := []float64{1.23, 2.32, 4.53, 7.23, 1.22, 1.11}

	t.Run("Test_fooA", func(t *testing.T) {
		a, b, err := fooA(arr)
		if err != nil {
			t.Errorf("fooA() error: %s", err)
		}
		if a != 7.23 || b != 4.53 {
			t.Errorf("fooA() sort arr fail.")
		}
	})
}

func Test_fooB(t *testing.T) {
	arr := []float64{1.23, 2.32, 4.53, 7.23, 1.22, 1.11}

	t.Run("Test_fooB", func(t *testing.T) {
		a, b, err := fooB(arr)
		if err != nil {
			t.Errorf("fooB() error: %s", err)
		}
		if a != 7.23 || b != 4.53 {
			t.Errorf("fooB() sort arr fail.")
		}
	})
}

func Test_fooC(t *testing.T) {
	arr := []float64{1.23, 2.32, 4.53, 7.23, 1.22, 1.11}

	t.Run("Test_fooC", func(t *testing.T) {
		a, b, err := fooC(arr)
		if err != nil {
			t.Errorf("fooC() error: %s", err)
		}
		if a != 7.23 || b != 4.53 {
			t.Errorf("fooC() sort arr fail.")
		}
	})
}
