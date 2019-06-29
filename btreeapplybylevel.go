package piscine

import "reflect"

func printOrderTraversal(root *TreeNode, niveau int, fn interface{})  {
        if root == nil { 
                return 
        }
        if niveau == 1 {
		Invoke(fn, root.Data)
        }else if niveau > 1 { 
                printOrderTraversal(root.Left,niveau - 1, fn)
                printOrderTraversal(root.Right,niveau - 1, fn)
        }
}

func BTreeApplyByLevel(root *TreeNode, fn interface{}) {
        h := BTreeLevelCount(root)
        for i := 1; i <= h; i++ { 
                printOrderTraversal(root, i,fn)
        }
}

func Invoke(fn interface{}, args ...string) {
    v := reflect.ValueOf(fn)
    rargs := make([]reflect.Value, len(args))
    for i, a := range args {
        rargs[i] = reflect.ValueOf(a)
    }
    v.Call(rargs)
}

