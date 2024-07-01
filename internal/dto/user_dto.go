package dto

import (
	"matchlove-services/internal/constant"
)

type (
	UserDeliveryAddressDTO struct {
		PhoneNumber  string               `json:"phoneNumber" validate:"required"`
		AddressName  string               `json:"addressName" validate:"required"`
		StreetName   string               `json:"streetName"`
		ProvinceName string               `json:"provinceName" validate:"required"`
		ProvinceCode string               `json:"provinceCode" validate:"required"`
		CityName     string               `json:"cityName" validate:"required"`
		CityCode     string               `json:"cityCode" validate:"required"`
		DistrictName string               `json:"districtName" validate:"required"`
		DistrictCode string               `json:"districtCode" validate:"required"`
		VillageName  string               `json:"villageName" validate:"required"`
		VillageCode  string               `json:"villageCode" validate:"required"`
		RT           string               `json:"rt" validate:"required"`
		RW           string               `json:"rw" validate:"required"`
		PostalCode   string               `json:"postalCode" validate:"required"`
		DetailOthers string               `json:"detailOthers" validate:"required"`
		Selected     int16                `json:"selected"`
		AddressType  constant.AddressType `json:"addressType" validate:"required"`
	}

	UserCartDTO struct {
		CarttID  string `json:"cartId"`
		Merchant struct {
			Product []ProductDTO `json:"products"`
		} `json:"merchant"`
		Quantity   int     `json:"qty"`
		TotalPrice float64 `json:"totalPrice"`
	}
)
