package model

type Space struct {
	W float32
	H float32
}

func (this Space) Constraint(vec Vec) Vec {
	return vec
}
