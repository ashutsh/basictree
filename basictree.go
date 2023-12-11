package basictree

import "fmt"

// So that you can create a node for the slice of nods needed to create the tree
// Level is the depth level at which the Cntent will be shown
func NewNode(level int, content string) Node {

	return Node{
		Level:   level,
		Content: content,
	}
}

type Node struct {
	Level     int
	graphics  string
	graphics1 string
	Content   string
}

// ├ ─ └ │
func (n *Node) levelPipe(c rune, l, h, offset int) {
	if n.Level != 0 {
		var gr []rune
		var temp []rune
		if n.graphics == "" {
			gr = make([]rune, (n.Level * h))
			temp = make([]rune, (n.Level * h))
		} else {
			gr = []rune(n.graphics)
			temp = []rune(n.graphics1)
		}
		id := (h * l) - (h - offset)
		gr[id] = c
		temp[id] = '│'
		temp[len(temp)-1] = '\n'
		if c != '│' {
			for j := 1; j < h-offset; j++ {
				gr[id+j] = '─'
			}
		}
		for i, r := range gr {
			if r == 0 {
				gr[i] = ' '
			} else {
				gr[i] = r
			}
		}
		for i, r := range temp {
			if r == 0 {
				temp[i] = ' '
			} else {
				temp[i] = r
			}
		}

		n.graphics1 = string(temp)
		n.graphics = string(gr)
	} else {
		n.graphics = ""
		
	}
}

// Creates a basic tree as of right now, the []Node is a basic node slice which contains only level int and content for the tree node.
// h is the horizontal spread of the tree. Starts from 1 and goes upward (default 4 if unacceptable value occured)
// v is the vertical spread of the tree. Starts from 0 and goes upwards (default 0 if unacceptable value occure)
// offset is the horizontal offset of the tree pipe with respect to the content, it should start from 0 but be less then h
func Tree(t []Node, h, v, offset int) {
	sisters := make(map[int][][]int)
	for level := 0; ; level++ {
		isMatched := false
		sister := make([]int, 0)
		streak := false
		for ti, r := range t {
			if r.Level == level {
				isMatched = true
				if streak {
					sister = append(sister, ti)
				} else {
					streak = true
					if len(sister) > 0 {
						sisters[level] = append(sisters[level], sister)
						sister = nil
					}
					sister = append(sister, ti)
				}

			}
			if r.Level < level && streak && len(sister) > 0 {
				streak = false
				sisters[level] = append(sisters[level], sister)
				sister = nil
			}
		}
		if len(sister) > 0 {
			sisters[level] = append(sisters[level], sister)
		}

		// To end the cycle if we have crossed the deepest level
		if !isMatched {
			break
		}
	}

	for levels, ranges := range sisters {
		for _, r := range ranges {
			if len(r) > 1 {
				for i := r[0]; i < r[len(r)-1]; i++ {
					t[i].levelPipe('│', levels, h, offset)
				}
				for ll, i := range r {
					if ll == len(r)-1 {
						t[i].levelPipe('└', levels, h, offset)
					} else {
						t[i].levelPipe('├', levels, h, offset)
					}
				}
			} else if len(r) == 1 {
				t[r[0]].levelPipe('└', levels, h, offset)
			}
		}
	}

	// For Printing the tree
	for _, r := range t {

		// For vertical height
		graphics := ""
		for j := 1; j <= v; j++ {
			graphics += r.graphics1
		}

		// Just for printing each node
		fmt.Printf("%s%s%s\n", graphics, r.graphics, r.Content)
	}

}
