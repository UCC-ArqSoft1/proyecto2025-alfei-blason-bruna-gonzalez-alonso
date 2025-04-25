package services

import "ejemplos/domain"

func GetHotelById(Id int) domain.Hotel {
	hotel := domain.Hotel{
		Id:         Id,
		Name:       "Holiday Inn",
		Roomscount: 20,
		Rating:     5,
		Location: domain.Location{
			Country: "Argentina",
			City:    "cordoba",
			Street:  "av. colon",
			Number:  100,
			Zipcode: 5000,
		},
	}
	return hotel
}
