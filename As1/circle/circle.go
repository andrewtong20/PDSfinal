/*
 This struct takes the user's inputted radius and finds the area,
 circumference if it is a circle and the volume and surface area if it is a
 sphere.
 */
package circle

import (
  "fmt"
  "math"
)


type circle struct {
  x,y,radius float64
}

func (c circle) getArea() float64 {
  return math.Pi*math.pow(c.radius,2)
}

func (c circle) getCircum() float64 {
  return 2*math.Pi*c.radius
}

func (c circle) getVolume() float64 {
  return 4*math.Pi*math.pow(c.radius,3)/3
}

func (c circle) getSurfaceA() float64 {
  return 4*math.Pi*math.pow(c.radius,2)
}
