package learner

type Wire struct {
	wid int32
	wtype string
	length float64
	attenuation float64
}

type Node struct {
	id int32
	retention float64
	depression float64
	excitation float64
}
