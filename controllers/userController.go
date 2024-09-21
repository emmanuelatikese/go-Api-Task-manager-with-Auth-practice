package controllers

import (
	"encoding/json"
	"net/http"
	"api.task/db"
	jwtFunc "api.task/jwt"
	"api.task/models"
	"api.task/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp (w http.ResponseWriter, r *http.Request) {
	var userModel routes.UserModel
	mongDb := apiDB.MongoDB
	userCollection := mongDb.Collection("userCollection")
	ctx := apiDB.Ctx

	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		apiUtils.JsonResponse(err.Error(), w, http.StatusBadRequest)
	}

	if userModel.Username == "" || string(userModel.Password) != string(userModel.Confirm_Password) {
		apiUtils.JsonResponse("Username invalid or Passwords mismatch", w, http.StatusNotAcceptable)
		return
	}
	if len(string(userModel.Password)) < 6 {
		apiUtils.JsonResponse("Password must be 6 or more", w, http.StatusNotAcceptable)
		return
	}

	HashPassword, err := bcrypt.GenerateFromPassword(userModel.Password, bcrypt.DefaultCost)
	if err != nil {
		apiUtils.JsonResponse(err.Error(), w, 500)
		return
	}


	insertId, err := userCollection.InsertOne(ctx, bson.M{
		"username": userModel.Username,
		"email": userModel.Email,
		"password": HashPassword,
	})
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
	}
	jwtFunc.GenerateToken(insertId.InsertedID, w)

	response := map[string]interface{}{
		"_id": insertId.InsertedID,
		"username": userModel.Username,
		"Email": userModel.Email,
	}

	apiUtils.JsonResponse(response, w, 200)
}