package resto

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

type foodCount struct {
	id    int
	count int
}

func appendIfNotExists(slice []int, s int) ([]int, error) {
	for _, v := range slice {
		if v == s {
			return slice, errors.New("Duplicate food menu id " + strconv.Itoa(v))
		}
	}
	return append(slice, s), nil
}

func RestoLog(file *os.File) ([]foodCount, error) {
	// Create a map to store the count of each menu item
	menuCounts := make(map[int]int)
	// Create a map to check the unique food items by eater id
	order := make(map[int][]int)
	// Read the log file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into eater_id and foodmenu_id
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			return nil, errors.New("Error: invalid log entry: " + line)
		}
		eaterID, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, errors.New("Error: invalid log entry: " + fields[0])
		}
		menuID, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, errors.New("Error: invalid menu ID: " + fields[1])
		}

		order[eaterID], err = appendIfNotExists(order[eaterID], menuID)
		if err != nil {
			return nil, err
		}
		// Update the count for this menu item
		menuCounts[menuID]++
	}

	// Sort the menu items by count in descending order
	counts := make([]foodCount, 0, len(menuCounts))
	for id, count := range menuCounts {
		counts = append(counts, foodCount{id, count})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	return counts[:3], nil
}
