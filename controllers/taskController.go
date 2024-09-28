package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	apiDB "api.task/db"
	model "api.task/models"
	apiUtils "api.task/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := apiDB.Ctx
	taskCollection := apiDB.TaskCollection

	var newtask model.TaskModel

	err := json.NewDecoder(r.Body).Decode(&newtask)
	if err != nil {
		apiUtils.JsonResponse(err, w, http.StatusBadRequest)
		return
	}

	insertId, err := taskCollection.InsertOne(ctx, newtask)
	if err != nil {
		apiUtils.JsonResponse(err, w, http.StatusNotAcceptable)
		return
	}

	response := &map[string]interface{}{
		"_id": insertId.InsertedID,
		"title": newtask.Title,
		"description":newtask.Description,
		"created_at": newtask.CreatedAt,
		"updated_at": newtask.UpdatedAt,
		"completed": newtask.Completed,
	}

	apiUtils.JsonResponse(response, w, 201)
}

func FindTask(w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	ctx := apiDB.Ctx
	taskCollection := apiDB.TaskCollection

	if id == "" {
		apiUtils.JsonResponse("Not found", w, 404)
		return
	}
	priId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}
	var filterTask bson.M // this is done for consistent response instead of routes.TaskModel type
	err = taskCollection.FindOne(ctx, bson.M{"_id": priId}).Decode(&filterTask)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}

	apiUtils.JsonResponse(filterTask, w, 200)
}


func FindAll (w http.ResponseWriter, r *http.Request){
	ctx := apiDB.Ctx
	taskCollection := apiDB.TaskCollection

	allTaskCursor, err := taskCollection.Find(ctx, bson.D{})
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}
	var alltask []bson.M // forgot that it was a list of bson.M
	err = allTaskCursor.All(ctx, &alltask) // forgot to bring &alltask
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}
	if alltask == nil {
		apiUtils.JsonResponse("No tasks available", w, 200)
		return
	}
	apiUtils.JsonResponse(alltask, w, 200)
}


func UpdateTask (w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]

	if id == ""{
		apiUtils.JsonResponse("not found", w, 404)
		return
	}
	var updateTask model.TaskModel
	ctx, taskCollection := apiDB.Ctx, apiDB.TaskCollection
	
	priId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}

	updateTask.UpdatedAt = time.Now()

	var filter bson.M
	err = taskCollection.FindOneAndUpdate(ctx,
		bson.M{"_id": priId}, bson.M{"$set": updateTask},
	options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&filter)
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}
	apiUtils.JsonResponse(filter, w, 201)
}

func DeleteOne (w http.ResponseWriter, r *http.Request){
	id := mux.Vars(r)["id"]
	if id == ""{
		apiUtils.JsonResponse("not found", w, 404)
		return
	}

	ctx, taskCollection := apiDB.Ctx, apiDB.TaskCollection

	deleteResult, err := taskCollection.DeleteOne(ctx, bson.D{{}})
	if err != nil {
		apiUtils.JsonResponse(err, w, http.StatusInternalServerError)
		return
	}

	if deleteResult.DeletedCount == 0 {
		apiUtils.JsonResponse("Tasks don't exit", w, 404)
		return
	}
	apiUtils.JsonResponse("Deleted successfully", w, 201)
}

func DeleteAll (w http.ResponseWriter, r *http.Request){
	ctx, taskCollection := apiDB.Ctx, apiDB.TaskCollection
	alltask, err := taskCollection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		apiUtils.JsonResponse(err, w, 500)
		return
	}
	if alltask.DeletedCount == 0 {
		apiUtils.JsonResponse("No tasks available", w, 404)
		return
	}
	apiUtils.JsonResponse("All deleted successfully", w, 201)
}