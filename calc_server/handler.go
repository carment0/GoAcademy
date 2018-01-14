// HTTP Handlers in Go is like Rail's Routes and Controllers
package main

import (
  "fmt"
  "net/http"
  "strconv"
)

const (
  ADD = "ADD"
  SUB = "SUBTRACT"
  DIV = "DIVIDE"
  MUL = "MULTIPLY"
)

// r = Request object
// w = what we will use to write the response and send it back to the client
// the endpoint expects two parameters = leftOp, rightOp e.g. 1 + 1 => 1, 1
func HandleAdd(w http.ResponseWriter, r *http.Request) {
  // We need to extract the Op from the request params e.g. lop, rop
  // localhost:3000/add?lop=1rop=1
  // the values come back as a string so we need to convert it
  // .PerseFloat takes (string, bit size)
  leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
  rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

  if leftErr == nil && rightErr == nil {
    result := fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp + rightOp)
    w.WriteHeader(http.StatusOK)
    // change the string into byte to send
    w.Write([]byte(result))
  } else {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Invalid query parameters! lop and rop should be float integers!"))

  }
}

// Since the handler functions are mostly the same for each calculation handler, we are going to create a Higher Order Function (a fn returning a fn)!
// the expected output is a http.HandleFunca
func OperationHandlerCreator(opType string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    leftOp, leftErr := strconv.ParseFloat(r.URL.Query().Get("lop"), 64)
    rightOp, rightErr := strconv.ParseFloat(r.URL.Query().Get("rop"), 64)

    if leftErr != nil || rightErr != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("Invalid query parameters! lop and rop should be float integers!"))
      return
    }

    var result string
    switch opType {
    case ADD:
      // .Sprintf = string print format
      result = fmt.Sprintf("%v + %v = %v", leftOp, rightOp, leftOp + rightOp)
    case SUB:
      result = fmt.Sprintf("%v - %v = %v", leftOp, rightOp, leftOp - rightOp)
    case DIV:
      result = fmt.Sprintf("%v / %v = %v", leftOp, rightOp, leftOp / rightOp)
    case MUL:
      result = fmt.Sprintf("%v * %v = %v", leftOp, rightOp, leftOp * rightOp)
    default:
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte("Something is wrong with the server"))
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(result))
  }
}
