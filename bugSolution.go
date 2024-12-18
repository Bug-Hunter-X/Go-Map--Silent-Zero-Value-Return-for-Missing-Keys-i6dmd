func main() {
    var m = make(map[int]int)
    m[1] = 1
    val, ok := m[2]
    if ok {
        fmt.Println("Value:", val)
    } else {
        fmt.Println("Key not found") // Handle the case where the key does not exist
    }
}    