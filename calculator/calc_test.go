package calculator

import "testing"

func TestAddTableDriven(t *testing.T) {
	tests := []struct { // Define a struct for each test case and create a slice of them
		name string
		a, b int
		want int
	}{
		{"both positive", 2, 3, 5},
		{"positive + zero", 5, 0, 5},
		{"negative + positive", -1, 4, 3},
		{"both negative", -2, -3, -5},
	}

	for _, tt := range tests { // Loop over each test case
		t.Run(tt.name, func(t *testing.T) { // Run each case as a subtest
			got := Add(tt.a, tt.b)
			if got != tt.want { // Check the result
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want) // Report failure if it doesn't match
			}
		})
	}
}

func TestDivideTableDriven(t *testing.T) {
	tests := []struct { // Define test cases
		name      string
		a, b      int
		want      int
		wantPanic bool
	}{
		{"normal division", 10, 2, 5, false},
		{"division by zero", 10, 0, 0, true},
	}

	for _, tt := range tests { // Loop over
		t.Run(tt.name, func(t *testing.T) { // Run subtest
			if tt.wantPanic { // Check for expected panic
				defer func() { // Recover from panic
					if r := recover(); r == nil {
						t.Errorf("Divide(%d, %d) did not panic", tt.a, tt.b)
					}
				}()
			}
			got := Divide(tt.a, tt.b) // Tests that do not panic
			if !tt.wantPanic && got != tt.want {
				t.Errorf("Divide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
