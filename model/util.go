package model

import (
	"strings"
	"fmt"
)

type ChanStruct struct {
	Processed int
	Unknowns []string
}

type Counter struct {
	maleCount   int
	femaleCount int
	unknown     int
}

func (c *Counter) Init() {
	c.femaleCount = 0
	c.maleCount = 0
	c.unknown = 0
}

func (c *Counter) AddGender(gender string) {
	if strings.ToLower(gender) == "male" {
		c.maleCount = c.maleCount + 1
		return
	}
	if strings.ToLower(gender) == "female" {
		c.femaleCount = c.femaleCount + 1
		return
	}
	c.unknown = c.unknown + 1
}

func (c *Counter) Print() {
	fmt.Printf("Total number of authors: %d\n\tMale: %d\n\tFemale: %d\n\tUnknown: %d\n",
		c.maleCount+c.femaleCount+c.unknown,
		c.maleCount,
		c.femaleCount,
		c.unknown)
}