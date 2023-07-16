package main

type HittableList struct {
	objects []Hittable
}

func (hitList *HittableList) Clear() {
	hitList.objects = nil
}

func (hitList *HittableList) Add(object Hittable) {
	hitList.objects = append(hitList.objects, object)
}

func (hitList HittableList) Hit(r Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	var tempRec HitRecord
	hitAnything := false
	closestSoFar := tMax

	for _, object := range hitList.objects {
		// var tempRec HitRecord
		if object.Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			*rec = tempRec
		}
	}

	return hitAnything
}
