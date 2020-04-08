/*
 This struct takes the user's inputted radius and finds the area,
 circumference if it is a circle and the volume and surface area if it is a
 sphere.
 */
package main

import (
  "math"
)


type Circle struct {
  x,y,radius float64
}
/*
//constructor
func New(x, y, radius float64) circle {
  c:= circle {x,y,radius}
  return c
}*/

func (c Circle) getArea() float64 {
  return math.Pi*math.Pow(c.radius,2)
}

func (c Circle) getCircum() float64 {
  return 2*math.Pi*c.radius
}

func (c Circle) getVolume() float64 {
  return 4*math.Pi*math.Pow(c.radius,3)/3
}

func (c Circle) getSurfaceA() float64 {
  return 4*math.Pi*math.Pow(c.radius,2)
}
