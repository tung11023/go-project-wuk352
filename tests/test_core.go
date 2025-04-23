package utils

import "testing"

func TestProcessData(t *testing.T) {
    data := ProcessData("dummy_path")
    if len(data) != 100 {
        t.Errorf("Expected 100 records, got %d", len(data))
    }
}
