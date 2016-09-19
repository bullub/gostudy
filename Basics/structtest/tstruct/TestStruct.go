package tstruct

type A struct {
	point *Point
	A *[]int
	F func()
}



func (a *A) GetPosition() *Point {
	if(nil == a.point) {
		a.point = new(Point);
		a.point.X = 1;
		a.point.Y = 2;
	}
	return a.point;
}
