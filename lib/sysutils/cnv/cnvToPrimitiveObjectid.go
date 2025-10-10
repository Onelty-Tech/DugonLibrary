package cnv

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
Convierte una interface a un primitive.ObjectID
*/
func ConvertPrimitiveObjectID(val any) (primitive.ObjectID, error) {
	if Id, ok := val.(primitive.ObjectID); ok {
		return Id, nil
	}
	return primitive.ObjectID{}, fmt.Errorf("error ConvertPrimitiveObjectID: Invalid data type, expected type 'primitive.ObjectID' found '%T'", val)
}
