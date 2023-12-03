package service

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func NewObject(w http.ResponseWriter, r *http.Request) {
	// client inputs
	type got struct {
		ObjectName string    `json:"objectName"`
		ObjectId   uuid.UUID `json:"objectId"`
	}

	var g got
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	object := Object{
		ObjectName: g.ObjectName,
		ObjectId:   g.ObjectId,
	}

	err := HandleNewObject(object)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	// client inputs
	type got struct {
		ObjectId uuid.UUID `json:"objectId"`
	}

	var g got
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := HandleDeleteObject(g.ObjectId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewRelation(w http.ResponseWriter, r *http.Request) {
	// client inputs
	type got struct {
		RelationType string    `json:"relation_type"`
		Object1Name  string    `json:"object1_name"`
		Object1Id    uuid.UUID `json:"object1_id"`
		Object2Name  string    `json:"object2_name"`
		Object2Id    uuid.UUID `json:"object2_id"`
	}

	var g got
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	relation := Relation{
		Type: g.RelationType,
		Object1: Object{
			ObjectName: g.Object1Name,
			ObjectId:   g.Object1Id,
		},
		Object2: Object{
			ObjectName: g.Object2Name,
			ObjectId:   g.Object2Id,
		},
	}

	err := HandleNewRelation(relation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	// Client inputs
	type got struct {
		ObjectId1    uuid.UUID `json:"objectId1"`
		ObjectId2    uuid.UUID `json:"objectId2"`
		RelationType string    `json:"relationType"`
	}

	var g got
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := HandleDeleteRelation(g.ObjectId1, g.ObjectId2, g.RelationType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CheckRelationExistence(w http.ResponseWriter, r *http.Request) {
	// client inputs
	type got struct {
		RelationType string    `json:"relation_type"`
		Object1Id    uuid.UUID `json:"object1_id"`
		Object2Id    uuid.UUID `json:"object2_id"`
	}

	var g got
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	relation := Relation{
		Type: g.RelationType,
		Object1: Object{
			ObjectId: g.Object1Id,
		},
		Object2: Object{
			ObjectId: g.Object2Id,
		},
	}

	exists, err := HandleRelationExists(relation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		Exists bool `json:"exists"`
	}{
		Exists: exists,
	}

	OkResponse(w, res)
}

func CheckRelationsExistence(w http.ResponseWriter, r *http.Request) {
	// client inputs
	type _relation struct {
		RelationType string    `json:"relation_type"`
		Object1Id    uuid.UUID `json:"object1_id"`
		Object2Id    uuid.UUID `json:"object2_id"`
	}

	type _relations struct {
		Relations []_relation `json:"relations"`
	}

	var got _relations
	if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	relations := make([]Relation, len(got.Relations))
	for i, r := range got.Relations {
		relations[i] = Relation{
			Type: r.RelationType,
			Object1: Object{
				ObjectId: r.Object1Id,
			},
			Object2: Object{
				ObjectId: r.Object2Id,
			},
		}
	}

	exist, err := HandleRelationsExist(relations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		Exist bool `json:"exist"`
	}{
		Exist: exist,
	}

	OkResponse(w, res)
}
