package domain

import "universities/api/universities"

type Panorama struct {
	UniversityId  string `json:"universityId"`
	Address       string `json:"address"`
	Name          string `json:"name"`
	FirstLocation string `json:"firstLocation"`
	LastLocation  string `json:"lastLocation"`
	Type          string `json:"type"`
}

func ConvertPanoramaTypeFromGrpc(t universities.PanoramaTypes) string {
	switch t {
	case universities.PanoramaTypes_Buildings:
		return "Корпуса"
	case universities.PanoramaTypes_Dormitories:
		return "Общежития"
	case universities.PanoramaTypes_Canteens:
		return "Столовые"
	case universities.PanoramaTypes_Other:
		return "Прочее"
	}

	return "Прочее"
}

func ConvertPanoramaToGrpc(t string) universities.PanoramaTypes {
	switch t {
	case "Корпуса":
		return universities.PanoramaTypes_Buildings
	case "Общежития":
		return universities.PanoramaTypes_Dormitories
	case "Столовые":
		return universities.PanoramaTypes_Canteens
	case "Прочее":
		return universities.PanoramaTypes_Other
	}

	return universities.PanoramaTypes_Other
}
