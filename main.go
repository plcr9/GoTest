package main

func main() {

}

func divide(x, y float32) (float 32, error) {
  var result float32

  if y == 0 {
    return result, errors.New("cannot divide by 0")
  }

  result = x / y
  return result
}
