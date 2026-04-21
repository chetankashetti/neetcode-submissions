func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    smap, tmap := make(map[rune]int), make(map[rune]int)
    for _, c := range s {
        smap[c]++
    }
    for _, c := range t {
        tmap[c]++
    }
    for _, c := range s {
        if smap[c] != tmap[c] {
            return false
        }
    }
    return true
}
