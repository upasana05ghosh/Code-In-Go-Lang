//https://leetcode.com/problems/dota2-senate/description/
func predictPartyVictory(senate string) string {
	r := []int{}
	d := []int{}

	for i, v := range senate {
		if v == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}

	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}

		r = r[1:] //pop
		d = d[1:] //pop
	}

	if len(r) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}