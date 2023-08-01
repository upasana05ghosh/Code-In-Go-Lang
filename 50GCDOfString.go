https://leetcode.com/problems/greatest-common-divisor-of-strings/description
func gcdOfStrings(s1 string, s2 string) string {
    //this checks if they've same char and something is common
    if(s1 + s2) != (s2 + s1) {
        return ""
    }

    //now we just need to find gcd of length
    gcdVal := gcd(len(s1), len(s2))
    return s1[: gcdVal];
}

func gcd(p, q int) int {
    if q == 0 {
        return p
    }

    return gcd(q, p%q)
}