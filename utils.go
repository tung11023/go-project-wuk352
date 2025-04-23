package utils

func ProcessData(path string) []map[string]int {
    // Simulate data processing
    var data []map[string]int
    for i := 0; i < 100; i++ {
        data = append(data, map[string]int{"id": i})
    }
    return data
}
