package calc_test

import (
    "testing"
    "github.com/lvreynoso/testintr"
)

func TestCalculate(t *testing.T) {
    if testintr.Calculate(2) != 4 {
        t.Error("Expected 2 + 2 to equal 4")
    }
}

func TestTableCalculate(t *testing.T) {
    var tests = []struct {
        input    int
        expected int
    }{
        {2, 4},
        {-1, 1},
        {0, 2},
        {-5, -3},
        {99999, 100001},
    }

    for _, test := range tests {
        if output := testintr.Calculate(test.input); output != test.expected {
            t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
        }
    }
}

// func BenchmarkCalculate(b *testing.B) {
//     for i := 0; i < b.N; i++ {
//         Calculate(2)
//     }
// }

func benchmarkCalculate(input int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        testintr.Calculate(input)
    }
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
