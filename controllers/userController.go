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
		apiUtils.JsonResponse("Username or Passwords invalid", w, http.StatusNotAcceptable)
		return

	}
	if len(userMap["password"]) < 6 {
		apiUtils.JsonResponse("Password must be 6 or more", w, http.StatusNotAcceptable)
		return
	}

	log.Print(userMap["password"])

	HashPassword, err := bcrypt.GenerateFromPassword([]byte(userMap["password"]), bcrypt.DefaultCost)
	if err != nil {
		apiUtils.JsonResponse(err.Error(), w, 500)
		return
	}
	userModel = model.UserModel{
		Username: userMap["username"],
		Email: userMap["email"],
		Password: HashPassword,
	}



	insertId, err := userCollection.InsertOne(ctx, userModel)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
	}
	jwtFunc.GenerateToken(insertId.InsertedID, w)

	response := map[string]interface{}{
		"_id": userModel.ID,
		"username": userModel.Username,
		"email": userModel.Email,
	}

	apiUtils.JsonResponse(response, w, 200)
}



func Login (w http.ResponseWriter, r *http.Request) {
	loginMap := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&loginMap)
	userCollection := apiDB.UserCollection
	ctx := apiDB.Ctx
	if err != nil {
		apiUtils.JsonResponse(err, w, http.StatusBadRequest)
		return
	}
	if loginMap["username"] == "" || loginMap["password"] == "" {
		apiUtils.JsonResponse("fill all fields", w, http.StatusNotAcceptable)
		return
	}

	var filterUser model.UserModel

	err = userCollection.FindOne(ctx, bson.M{"username": loginMap["username"]}).Decode(filterUser)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			apiUtils.JsonResponse("User and password don't exist", w, http.StatusNotFound)
			return
		}
	}
	strpassword := string(filterUser.Password)
	log.Print(strpassword)
	log.Print(loginMap["password"])
	err = bcrypt.CompareHashAndPassword(filterUser.Password, []byte(loginMap["password"]))
	if err != nil {
		apiUtils.JsonResponse("Wrong password", w, http.StatusNotAcceptable)
		return
	}

	jwtFunc.GenerateToken(filterUser.ID, w)
	apiUtils.JsonResponse("Login in successfully", w, 200)
}

