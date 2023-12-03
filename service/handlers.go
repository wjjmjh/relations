package service

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var Driver neo4j.Driver

func HandleNewObject(object Object) error {
	session := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(
			"CREATE (n:Object { objectName: $objectName, objectId: $objectId })",
			map[string]interface{}{
				"objectName": object.ObjectName,
				"objectId":   object.ObjectId.String(),
			},
		)
		return nil, err
	})
	if err != nil {
		return err
	}

	return nil
}

func HandleDeleteObject(objectId uuid.UUID) error {
	session := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run(
			"MATCH (n:Object { objectId: $objectId }) DELETE n",
			map[string]interface{}{
				"objectId": objectId.String(),
			},
		)
		return nil, err
	})

	if err != nil {
		return err
	}

	return nil
}

func HandleNewRelation(relation Relation) error {
	session := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (a:Object), (b:Object) WHERE a.objectId = $objectId1 AND b.objectId = $objectId2 CREATE (a)-[r:" + relation.Type + "]->(b) RETURN r"

		_, err := tx.Run(query, map[string]interface{}{
			"objectId1": relation.Object1.ObjectId.String(),
			"objectId2": relation.Object2.ObjectId.String(),
		})
		return nil, err
	})
	if err != nil {
		return err
	}

	return nil
}

func HandleDeleteRelation(objectId1, objectId2 uuid.UUID, relationType string) error {
	session := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (a:Object)-[r:" + relationType + "]->(b:Object) WHERE a.objectId = $objectId1 AND b.objectId = $objectId2 DELETE r"

		_, err := tx.Run(query, map[string]interface{}{
			"objectId1": objectId1.String(),
			"objectId2": objectId2.String(),
		})
		return nil, err
	})

	if err != nil {
		return err
	}

	return nil
}

func HandleRelationExists(relation Relation) (bool, error) {
	session := Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run(
		"MATCH (a:Object {objectId: $objectId1})-[c:"+relation.Type+"]->(b:Object {objectId: $objectId2}) RETURN c",
		map[string]interface{}{
			"objectId1": relation.Object1.ObjectId,
			"objectId2": relation.Object2.ObjectId,
		},
	)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}

	if result.Next() {
		return true, nil
	}

	return false, nil
}

func HandleRelationsExist(relations []Relation) (bool, error) {
	for _, relation := range relations {
		exists, err := HandleRelationExists(relation)
		if err != nil {
			return false, err
		}
		if !exists {
			return false, nil
		}
	}

	return true, nil
}
