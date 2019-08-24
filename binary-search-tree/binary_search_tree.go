package binarysearchtree

type SearchTreeData struct {
	left *SearchTreeData
	data int
	right *SearchTreeData
}

func Bst(d int) (*SearchTreeData) {

	return &SearchTreeData{nil, d, nil}
}

func (bst *SearchTreeData) Insert(newData int) {

	// Don't really need to check if tree is null here, as its always
	// constructed before calling this function in the tests.

	if (newData <= bst.data) {
		// Is the left node empty? If so, then insert here
		if (bst.left == nil) {
			bst.left = Bst(newData)
		} else {
			// If left node already exists, call insert on that node
			bst.left.Insert(newData)
		}

	} else if (newData > bst.data) {
		// Same logic as the case above.
		if (bst.right == nil) {
			bst.right = Bst(newData)
		} else {
			bst.right.Insert(newData)
		}
	}
}


func (bst *SearchTreeData) MapString(fun func(int) string) (treeAsString []string) {
	if (bst.left != nil) {
		treeAsString = append(treeAsString, bst.left.MapString(fun) ...)
	}

	treeAsString = append(treeAsString, fun(bst.data))

	if (bst.right != nil) {
		treeAsString = append(treeAsString, bst.right.MapString(fun) ...)
	}

	return treeAsString
}

func (bst *SearchTreeData) MapInt(fun func(int) int) (treeAsInt []int) {
	if (bst.left != nil) {
		treeAsInt = append(treeAsInt, bst.left.MapInt(fun) ...)
	}

	treeAsInt = append(treeAsInt, fun(bst.data))

	if (bst.right != nil) {
		treeAsInt = append(treeAsInt, bst.right.MapInt(fun) ...)
	}

	return
}
