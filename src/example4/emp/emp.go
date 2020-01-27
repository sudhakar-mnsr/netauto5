package emp

type emp struct {
   Name string
   Sal float64
}

type Manager struct {
   emp
   Teamsize int
}
