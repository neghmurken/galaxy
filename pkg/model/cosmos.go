package model

import rl "github.com/gen2brain/raylib-go/raylib"

type Cosmos struct {
	Bodies []*Body
	Scale  float32
	Bounds Space
}

func MakeCosmos(scale float32, space Space) *Cosmos {
	return &Cosmos{Scale: scale, Bounds: space}
}

func (this *Cosmos) Add(body *Body) {
	this.Bodies = append(this.Bodies, body)
}

func (this *Cosmos) Remove(index int) {
	var s []*Body = this.Bodies
	lastIndex := len(this.Bodies) - 1
	s[index] = this.Bodies[lastIndex]
	this.Bodies = s[:lastIndex]
}

func (this *Cosmos) Update(dt float32) {
	for _, body := range this.Bodies {
		body.Grow(dt)
		body.ApplyForce(this.GatherForces(body), dt, this.Bounds)
	}

	for i, body := range this.Bodies {
		for j := i; j < len(this.Bodies); j++ {
			if body.Collides(this.Bodies[j]) {
				body.MeldWidth(this.Bodies[j])
				this.Remove(j)
				j++
			}
		}
	}
}

func (this *Cosmos) GatherForces(body *Body) Vec {
	resultant := rl.Vector2Zero()

	for _, other := range this.Bodies {
		if other != body {
			resultant = rl.Vector2Add(resultant, body.GravityFrom(other))
		}
	}

	return rl.Vector2Scale(resultant, this.Scale)
}
