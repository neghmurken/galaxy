package model

import (
	"math/rand"

	u "github.com/neghmurken/galaxy/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	EPSILON float32 = 0.5
)

type Vec = rl.Vector2

type Body struct {
	Pos  Vec
	Vel  Vec
	Mass float32
}

func NewBody(pos, vel Vec, mass float32) *Body {
	return &Body{pos, vel, mass}
}

func NewStaticBody(pos Vec, mass float32) *Body {
	return NewBody(pos, rl.Vector2Zero(), mass)
}

func (this *Body) MeldWidth(other *Body) {
	this.Vel = rl.Vector2Lerp(
		other.Vel,
		this.Vel,
		this.Mass/(this.Mass+other.Mass),
	)
	this.Mass += other.Mass
}

func (this *Body) Collides(other *Body) bool {
	if this == other {
		return false
	}

	return rl.Vector2Length(this.Distance(other)) <= EPSILON
}

func (this *Body) Distance(other *Body) Vec {
	return rl.Vector2Subtract(other.Pos, this.Pos)
}

func (this *Body) GravityFrom(other *Body) Vec {
	dist := this.Distance(other)
	len := max(rl.Vector2Length(dist), EPSILON)

	return rl.Vector2Scale(
		rl.Vector2Normalize(dist),
		u.Gravity(this.Mass, other.Mass, len),
	)
}

func (this *Body) ApplyForce(force Vec, dt float32) {
	prevVel := this.Vel
	a := rl.Vector2Scale(force, 1./this.Mass)

	this.Vel = rl.Vector2Add(prevVel, rl.Vector2Scale(a, dt))
	this.Pos = rl.Vector2Add(this.Pos, rl.Vector2Scale(prevVel, 0.5*dt))
}

func RandVec(w, h float32) Vec {
	return Vec{X: rand.Float32()*w - w/2, Y: rand.Float32()*h - h/2}
}
