package main

import "testing"

func TestMoonSystem_ApplyStep(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		//<x=-1, y=0,   z=2>
		//<x=2,  y=-10, z=-7>
		//<x=4,  y=-8,  z=8>
		//<x=3,  y=5,  z=-1>
		system := &MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{-1, 0, 2}},
			{position: ThreeDCoord{2, -10, -7}},
			{position: ThreeDCoord{4, -8, 8}},
			{position: ThreeDCoord{3, 5, -1}},
		}}
		system.ApplyStep()
		//pos=<x= 2, y=-1, z= 1>, vel=<x= 3, y=-1, z=-1>
		//	pos=<x= 3, y=-7, z=-4>, vel=<x= 1, y= 3, z= 3>
		//	pos=<x= 1, y=-7, z= 5>, vel=<x=-3, y= 1, z=-3>
		//	pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>
		expected1 := MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{2, -1, 1}, velocity: ThreeDCoord{3, -1, -1}},
			{position: ThreeDCoord{3, -7, -4}, velocity: ThreeDCoord{1, 3, 3}},
			{position: ThreeDCoord{1, -7, 5}, velocity: ThreeDCoord{-3, 1, -3}},
			{position: ThreeDCoord{2, 2, 0}, velocity: ThreeDCoord{-1, -3, 1}},
		}}
		compare(expected1, *system, t)

		system.ApplyStep()
		//After 2 steps:
		//  pos=<x= 5, y=-3, z=-1>, vel=<x= 3, y=-2, z=-2>
		//	pos=<x= 1, y=-2, z= 2>, vel=<x=-2, y= 5, z= 6>
		//	pos=<x= 1, y=-4, z=-1>, vel=<x= 0, y= 3, z=-6>
		//	pos=<x= 1, y=-4, z= 2>, vel=<x=-1, y=-6, z= 2>
		expected2 := MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{5, -3, -1}, velocity: ThreeDCoord{3, -2, -2}},
			{position: ThreeDCoord{1, -2, 2}, velocity: ThreeDCoord{-2, 5, 6}},
			{position: ThreeDCoord{1, -4, -1}, velocity: ThreeDCoord{0, 3, -6}},
			{position: ThreeDCoord{1, -4, 2}, velocity: ThreeDCoord{-1, -6, 2}},
		}}
		compare(expected2, *system, t)
	})
}

func TestSimulate(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		system := &MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{-1, 0, 2}},
			{position: ThreeDCoord{2, -10, -7}},
			{position: ThreeDCoord{4, -8, 8}},
			{position: ThreeDCoord{3, 5, -1}},
		}}
		Simulate(system, 10)
		//After 10 steps:
		//  pos=<x= 2, y= 1, z=-3>, vel=<x=-3, y=-2, z= 1>
		//	pos=<x= 1, y=-8, z= 0>, vel=<x=-1, y= 1, z= 3>
		//	pos=<x= 3, y=-6, z= 1>, vel=<x= 3, y= 2, z=-3>
		//	pos=<x= 2, y= 0, z= 4>, vel=<x= 1, y=-1, z=-1>
		expected := MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{2, 1, -3}, velocity: ThreeDCoord{-3, -2, 1}},
			{position: ThreeDCoord{1, -8, 0}, velocity: ThreeDCoord{-1, 1, 3}},
			{position: ThreeDCoord{3, -6, 1}, velocity: ThreeDCoord{3, 2, -3}},
			{position: ThreeDCoord{2, 0, 4}, velocity: ThreeDCoord{1, -1, -1}},
		}}
		compare(expected, *system, t)
	})
}

func compare(expected, actual MoonSystem, t *testing.T) {
	for i, eMoon := range expected.moons {
		result := actual.moons[i]
		if result.position != eMoon.position {
			t.Errorf("Moon %d position %v not equal to expected %v", i, actual.moons[i], eMoon)
		}
		if result.velocity != eMoon.velocity {
			t.Errorf("Moon %d velocity %v not equal to expected %v", i, actual.moons[i], eMoon)
		}
	}
}

func TestMoonSystem_TotalEnergy(t *testing.T) {
	ms := MoonSystem{moons: []*Moon{
		{position: ThreeDCoord{2, 1, -3}, velocity: ThreeDCoord{-3, -2, 1}},
		{position: ThreeDCoord{1, -8, 0}, velocity: ThreeDCoord{-1, 1, 3}},
		{position: ThreeDCoord{3, -6, 1}, velocity: ThreeDCoord{3, 2, -3}},
		{position: ThreeDCoord{2, 0, 4}, velocity: ThreeDCoord{1, -1, -1}},
	}}
	result := ms.TotalEnergy()
	if result != 179 {
		t.Errorf("Wrong total energy expected 179 got %d", result)
	}
}

func TestExamples(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		ms := ParseMoonSystem([]string{
			"<x=-1, y=0, z=2>",
			"<x=2, y=-10, z=-7>",
			"<x=4, y=-8, z=8>",
			"<x=3, y=5, z=-1>",
		})
		Simulate(ms, 10)
		compare(MoonSystem{moons: []*Moon{
			{position: ThreeDCoord{2, 1, -3}, velocity: ThreeDCoord{-3, -2, 1}},
			{position: ThreeDCoord{1, -8, 0}, velocity: ThreeDCoord{-1, 1, 3}},
			{position: ThreeDCoord{3, -6, 1}, velocity: ThreeDCoord{3, 2, -3}},
			{position: ThreeDCoord{2, 0, 4}, velocity: ThreeDCoord{1, -1, -1}},
		}}, *ms, t)
		result := ms.TotalEnergy()
		if result != 179 {
			t.Errorf("Wrong total energy expected 179 got %d", result)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		ms := ParseMoonSystem([]string{
			"<x=-8, y=-10, z=0>",
			"<x=5, y=5, z=10>",
			"<x=2, y=-7, z=3>",
			"<x=9, y=-8, z=-3>",
		})
		Simulate(ms, 100)
		result := ms.TotalEnergy()
		if result != 1940 {
			t.Errorf("Wrong total energy expected 1940 got %d", result)
		}
	})
}
