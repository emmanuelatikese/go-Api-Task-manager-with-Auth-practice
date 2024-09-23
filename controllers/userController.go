package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"api.task/db"
	jwtFunc "api.task/jwt"
	"api.task/models"
	"api.task/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)




func SignUp (w http.ResponseWriter, r *http.Request) {
	var userModel model.UserModel
	ctx := apiDB.Ctx
	userCollection := apiDB.UserCollection
	userMap := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&userMap)
	if err != nil {
		apiUtils.JsonResponse(err.Error(), w, http.StatusBadRequest)
		return
	}
	var commonUser model.UserModel

	err = userCollection.FindOne(ctx, bson.M{"username": userMap["username"]}).Decode(commonUser)
	if err != mongo.ErrNoDocuments{
		apiUtils.JsonResponse("Username used", w, http.StatusNotAcceptable)
		return
	}


	if userMap["username"] == "" || userMap["password"] != userMap["confirmed_password"] {
		apiUtils.JsonResponse("Username invalid or Passwords mismatch", w, http.StatusNotAcceptable)
		return

	}
	if len(userMap["password"]) < 6 {
		log.Print(len(userMap["password"]))
		apiUtils.JsonResponse("Password must be 6 or more", w, http.StatusNotAcceptable)
		return
	}

	userModel = model.UserModel{
		Username: userMap["username"],
		Email: userMap["email"],
		Password: []byte(userMap["password"]),
	}

	HashPassword, err := bcrypt.GenerateFromPassword(userModel.Password, bcrypt.DefaultCost)
	if err != nil {
		apiUtils.JsonResponse(err.Error(), w, 500)
		return
	}
	userModel.Password = HashPassword


	insertId, err := userCollection.InsertOne(ctx, userModel)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
	}
	jwtFunc.GenerateToken(insertId.InsertedID, w)

	response := map[string]interface{}{
		"_id": insertId.InsertedID,
		"username": userModel.Username,
		"email": userModel.Email,
	}

	apiUtils.JsonResponse(response, w, 200)
}


