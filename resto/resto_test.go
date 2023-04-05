package resto

import (
	"os"
	"testing"
)

func TestRestoLog(t *testing.T) {
	// Test case 1: valid log file with no errors
	file, err := os.Open("testdata/valid_log.txt")
	if err != nil {
		t.Errorf("Failed to open log file: %v", err)
	}
	defer file.Close()

	counts, err := RestoLog(file)
	expectedCounts := []foodCount{
		{
			id:    100,
			count: 2,
		},
		{
			id:    200,
			count: 2,
		},
		{
			id:    500,
			count: 1,
		},
	}

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(counts) != 3 {
		t.Errorf("Expected 3 menu items, got %d", len(counts))
	}
	for i := 0; i < 3; i++ {
		if expectedCounts[0].id != counts[0].id && expectedCounts[0].count != counts[0].count {
			t.Errorf("Unexpected error")
		}
		if expectedCounts[1].id != counts[1].id && expectedCounts[1].count != counts[1].count {
			t.Errorf("Unexpected error")
		}
		if expectedCounts[2].id != counts[2].id && expectedCounts[2].count != counts[2].count {
			t.Errorf("Unexpected error")
		}
	}

	// Test case 2: log file with invalid entry
	file, err = os.Open("testdata/invalid_entry_log.txt")
	if err != nil {
		t.Errorf("Failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = RestoLog(file)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 3: log file with invalid menu ID
	file, err = os.Open("testdata/invalid_menu_id_log.txt")
	if err != nil {
		t.Errorf("Failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = RestoLog(file)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	// Test case 4: log file with duplicate food menu ID for an eater ID
	file, err = os.Open("testdata/duplicate_food_menu_id_log.txt")
	if err != nil {
		t.Errorf("Failed to open log file: %v", err)
	}
	defer file.Close()

	_, err = RestoLog(file)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
