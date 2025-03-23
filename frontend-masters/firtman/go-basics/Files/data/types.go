package data

type distance float32 // miles
type distanceKm float32

// Method
func (miles distance) ToKm() distanceKm {
	return distanceKm(1.60934 * miles)
}

func (km distanceKm) ToMiles() distance {
	return distance(km / 1.60934)
}

func test() {
	d := distance(34.5)
	km := d.ToKm()
	km.ToMiles()
	print(d)
}
