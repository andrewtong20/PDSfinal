/*
front11:

Given 2 int arrays, a and b, of any length, return a new array with the
first element of each array. If either array is length 0, ignore that array.

*/


package main

func front11(inA []int, inB []int) []int {
  a:=inA
  b:=inB
  if(len(a) > 0 && len(b) > 0) {
      return []int{a[0], b[0]}
  } else if(len(a) > 0) {
      return []int{a[0]}
  } else if(len(b) > 0) {
      return []int{b[0]}
  }

  return make([]int,0)
}
