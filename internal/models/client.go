package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClientMg struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Surname string             `json:"surname" bson:"surname"`
	Age     int                `json:"age" bson:"age"`
	Phone   string             `json:"phone_number" bson:"phone_number"`
	Email   string             `json:"email" bson:"email"`
}

type Client struct {
	Id      string `json:"_id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Surname string `json:"surname" bson:"surname"`
	Age     int    `json:"age" bson:"age"`
	Phone   string `json:"phone_number" bson:"phone_number"`
	Email   string `json:"email" bson:"email"`
}

func MapClient(cl *ClientMg) *Client {
	return &Client{
		Id:      cl.Id.Hex(),
		Name:    cl.Name,
		Surname: cl.Surname,
		Age:     cl.Age,
		Phone:   cl.Phone,
		Email:   cl.Email,
	}
}
