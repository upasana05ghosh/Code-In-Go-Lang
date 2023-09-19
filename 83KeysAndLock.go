//https://leetcode.com/problems/keys-and-rooms/description/
func canVisitAllRooms(rooms [][]int) bool {
	set := make(map[int]bool)
	dfs(rooms, 0, set)
	set[0] = true
	fmt.Println(set, len(set), len(rooms)-1)
	return len(set) == len(rooms)
}

func dfs(rooms [][]int, index int, set map[int]bool) {
	for i := 0; i < len(rooms[index]); i++ {
		if !set[rooms[index][i]] {
			set[rooms[index][i]] = true
			dfs(rooms, rooms[index][i], set)
		}
	}
}