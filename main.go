package main

import (
  "fmt"
  //"github.com/stretchr/testify/assert"
)

type Neuron struct {
  weights []float64
  bias    float64
}

func main() {

	inputs := [4]float64{1, 2, 3, 2.5}

  n1 := Neuron{[]float64{0.2, 0.8, -0.5, 1.0}, 2.0}
  n2 := Neuron{[]float64{0.5, -0.91, 0.26, -0.5}, 3}
  n3 := Neuron{[]float64{-0.26, -0.27, 0.17, 0.87}, 0.5}

  neurons := []Neuron{n1 ,n2, n3}
  var outputs []float64

  for i := 0; i < len(neurons); i++ {
    neuron := neurons[i]
    output := 0.0 

    for k := 0; k < len(inputs); k++ {
      output += inputs[k] * neuron.weights[k]
    }
    output += neuron.bias

    outputs = append(outputs, output)
  } 

	fmt.Println(outputs)
}
