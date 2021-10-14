package algo

func SequenceExists(main, seq []int) bool {
	count := 1
	for i := 0; i < len(main)-len(seq)+1; i++ {
		if main[i] == seq[0] {
			for j := 1; j < len(seq); j++ {
				if main[i+j] == seq[j] {
					count += 1
				} else {
					count = 1
					break
				}
			}
		}
	}
	return count == len(seq)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Find(key int) bool {
	if tn == nil {
		return false
	}
	if tn.Val == key {
		return true
	}
	left := tn.Left.Find(key)
	if left {
		return true
	}
	right := tn.Right.Find(key)

	return right
}
