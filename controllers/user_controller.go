package controllers

import (
	"context"
	"fiber-api/configs"
	"fiber-api/models"
	"fiber-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Getting users model from db
var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// c in func argument is similar to request object in express
func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()},
			})
	}

	// create new user document
	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Location: user.Location,
	}

	// insert newUser document inside user collection
	result, err := userCollection.InsertOne(ctx, newUser)

	// check for error while inserting
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	// If no errors
	return c.Status(http.StatusCreated).JSON(
		responses.UserResponse{
			Status:  http.StatusCreated,
			Success: true,
			Message: "success",
			Data:    &fiber.Map{"data": result},
		})
}

// c in func argument is similar to request object in express
func GetUserById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var user models.User
	defer cancel()

	// objId is bson type id given by mongodb
	objId, _ := primitive.ObjectIDFromHex(userId)

	// checking if id exists = if user exists
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	// if no error, return the "Decoded" user
	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: "success",
			Data:    &fiber.Map{"data": user},
		})
}

func EditUserById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Params("userId")
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(
			responses.UserResponse{
				Status:  http.StatusBadRequest,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": validationErr.Error()},
			})
	}

	update := bson.M{"name": user.Name, "location": user.Location}

	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Success: false,
				Message: "error",
				Data:    &fiber.Map{"data": err.Error()},
			})
	}
	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Success: false,
					Message: "error",
					Data:    &fiber.Map{"data": err.Error()},
				})
		}
	}

	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{
			Status:  http.StatusOK,
			Success: false,
			Message: "success",
			Data:    &fiber.Map{"data": updatedUser},
		})
}
