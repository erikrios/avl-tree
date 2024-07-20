package main

type Imbalance int

const (
	LeftLeftImbalance Imbalance = iota
	RightRightImbalance
	LeftRightImbalance
	RightLeftImbalance
)
