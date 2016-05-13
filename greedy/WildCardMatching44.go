package main
	// 36ms
func isMatch(s string, p string) bool {
	i:=0
	j:=0
	star := -1
	s_star := 0
	s_len := len(s)
	p_len := len(p)
	for i<s_len{
		if i < s_len && j < p_len && (s[i] == p[j] || p[j] == '?'){
			i +=1
			j += 1
		} else if j< p_len && p[j] == '*'{
			star = j
			s_star = i
			j +=1
		} else if star != -1{
			j = star+1
			s_star += 1
			i = s_star
		} else{
			return false
		}
	}
	for j<p_len && p[j] == '*'{
		j+=1
	}
	return j==p_len
}

func main() {
	println(isMatch("aa343","a*3"))
}
