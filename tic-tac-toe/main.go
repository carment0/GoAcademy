package main


func main() {
  p1 := &HumanPlayer{"Momo", "X"}
  cp := &ComputerPlayer{"HAL9000", "O"}
  g := NewGame(p1, cp)
  g.Start()
}
