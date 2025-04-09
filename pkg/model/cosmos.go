package model

import rl "github.com/gen2brain/raylib-go/raylib"

type Cosmos struct {
	Bodies []*Body
	Scale  float32
}

func MakeCosmos(scale float32) *Cosmos {
	return &Cosmos{Scale: scale}
}

func (this *Cosmos) AddBody(body *Body) {
	this.Bodies = append(this.Bodies, body)
}

func (this *Cosmos) Update(dt float32) {
	for _, body := range this.Bodies {
		body.ApplyForce(this.GatherForces(body), dt)
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
