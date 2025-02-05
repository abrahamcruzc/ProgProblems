package main

func ZeroFuel(distanceToPump, mpg, fuelLeft int) bool {
	if mpg*fuelLeft >= distanceToPump {
		return true
	}

	return false
}
