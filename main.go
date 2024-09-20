package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Neuron struct {
  weights []float64
  bias    float64
}

func DotProduct(vector1 []float64, vector2 []float64) float64 {
  output := 0.0 

  //dot product between two vectors
  for k := 0; k < len(vector1); k++ {
    output += vector1[k] * vector2[k]
  }
  return output
}

func RunModel() {
  inputs := []float64{1, 2, 3, 2.5}

  n1 := Neuron{[]float64{0.2, 0.8, -0.5, 1.0}, 2.0}
  n2 := Neuron{[]float64{0.5, -0.91, 0.26, -0.5}, 3}
  n3 := Neuron{[]float64{-0.26, -0.27, 0.17, 0.87}, 0.5}

  neurons := []Neuron{n1 ,n2, n3}
  var outputs []float64

  for i := 0; i < len(neurons); i++ {
    neuron := neurons[i]
    output := DotProduct(inputs, neuron.weights) + neuron.bias
    outputs = append(outputs, output)
  } 

	fmt.Println(outputs)
}

func main() {

  a := mat.NewDense(1, 3, []float64{1,2,3})
  b := mat.NewDense(1, 3, []float64{2,3,4}).T()

  result := mat.NewDense(1, 1, nil)
  result.Product(a, b)
	
  fmt.Println(mat.Formatted(result))

}
