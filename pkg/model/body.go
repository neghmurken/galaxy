package model

import (
	"math/rand"

	u "github.com/neghmurken/galaxy/pkg/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	EPSILON   float32 = 0.5
	GROW_RATE float32 = 100
)

type Vec = rl.Vector2

type Body struct {
	Pos        Vec
	Vel        Vec
	Size       float32
	SizeGrowth float32
}

func NewBody(pos, vel Vec, size float32) *Body {
	return &Body{Pos: pos, Vel: vel, Size: size}
}

func NewStaticBody(pos Vec, mass float32) *Body {
	return NewBody(pos, rl.Vector2Zero(), mass)
}

func (this *Body) MeldWidth(other *Body) {
	factor := this.Size / (this.Size + other.Size)

	this.Vel = rl.Vector2Lerp(other.Vel, this.Vel, factor)
	this.Pos = rl.Vector2Lerp(other.Pos, this.Pos, factor)
	this.SizeGrowth += min(other.Size, this.Size)
	this.Size = max(this.Size, other.Size)
}

func (this *Body) GetMass() float32 {
	return u.SizeToMass(this.Size)
}

func (this *Body) GetKineticEnergy() float32 {
	return .5 * this.GetMass() * rl.Vector2LengthSqr(this.Vel)
}

func (this *Body) Collides(other *Body) bool {
	if this == other {
		return false
	}

	return rl.Vector2Length(this.Distance(other)) <= max(this.Size, other.Size)
}

func (this *Body) Distance(other *Body) Vec {
	return rl.Vector2Subtract(other.Pos, this.Pos)
}

func (this *Body) GravityFrom(other *Body) Vec {
	dist := this.Distance(other)
	len := max(rl.Vector2Length(dist), EPSILON)
	dir := rl.Vector2Normalize(dist)

	attraction := rl.Vector2Scale(
		dir,
		u.Gravity(this.GetMass(), other.GetMass(), len),
	)

	repulsion := rl.Vector2Scale(
		rl.Vector2Negate(dir),
		u.Explosion(
			this.GetMass(),
			u.SizeToMass(other.SizeGrowth),
			len,
			50000/len,
		),
	)

	return rl.Vector2Add(attraction, repulsion)
}

func (this *Body) ApplyForce(force Vec, dt float32, bounds Space) {
	prevVel := this.Vel
	a := rl.Vector2Scale(force, 1./this.GetMass())

	this.Vel = rl.Vector2Add(prevVel, rl.Vector2Scale(a, dt))
	this.Pos = bounds.Constraint(rl.Vector2Add(
		this.Pos,
		rl.Vector2Scale(prevVel, 0.5*dt),
	))
}

func (this *Body) Grow(dt float32) {
	if this.SizeGrowth != 0 {
		deltaGrowth := min(GROW_RATE*dt, this.SizeGrowth)
		this.Size += deltaGrowth
		this.SizeGrowth -= deltaGrowth
	}
}

func RandVec(w, h float32) Vec {
	return Vec{X: rand.Float32()*w - w/2, Y: rand.Float32()*h - h/2}
}
