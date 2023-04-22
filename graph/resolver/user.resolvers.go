package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	db "github.com/ghost-codes/uber/db/sqlc"
	"github.com/ghost-codes/uber/graph"
	"github.com/ghost-codes/uber/graph/model"
	"github.com/ghost-codes/uber/kafkaConfig"
	"github.com/golang/geo/s2"
	"github.com/segmentio/kafka-go"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, data model.CreateUserData) (*db.UserMetaData, error) {
	token, err := r.FirebaseAuth.VerifyIDToken(ctx, data.FirebaseAuthID)

	if err != nil {
		return nil, err
	}

	args := db.CreateUserMetaDataParams{
		ID:          token.UID,
		DateOfBirth: data.DateOfBirth,
		PhoneNumber: data.PhoneNumber,
	}

	metatData, err := r.Store.CreateUserMetaData(ctx, args)
	if err != nil {
		return nil, err
	}

	_, err = r.FirebaseAuth.CustomTokenWithClaims(ctx, metatData.ID, map[string]interface{}{"type": "client"})
	if err != nil {
		return nil, err
	}

	return &metatData, nil
}

// CreateSession is the resolver for the createSession field.
func (r *mutationResolver) CreateSession(ctx context.Context, tokenID string) (*model.Session, error) {
	token, err := r.FirebaseAuth.VerifyIDToken(ctx, tokenID)
	if err != nil {
		return nil, err
	}

	user, err := r.Store.FetchUserMetaDataByID(ctx, token.UID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		session := model.Session{
			IsSignupComplete: false,
		}

		return &session, nil
	}

	claims := map[string]interface{}{
		"userType": model.TypeClient,
	}

	err = r.FirebaseAuth.SetCustomUserClaims(ctx, token.UID, claims)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		IsSignupComplete: true,
		User:             &user,
	}

	return &session, nil
}

// UserMetaData is the resolver for the userMetaData field.
func (r *queryResolver) UserMetaData(ctx context.Context, id string) (*db.UserMetaData, error) {
	userMetaData, err := r.Store.FetchUserMetaDataByID(ctx, id)
	return &userMetaData, err
}

// Source is the resolver for the source field.
func (r *rideHistoryResolver) Source(ctx context.Context, obj *db.RideHistory) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: Source - source"))
}

// Destination is the resolver for the destination field.
func (r *rideHistoryResolver) Destination(ctx context.Context, obj *db.RideHistory) (*model.Location, error) {
	panic(fmt.Errorf("not implemented: Destination - destination"))
}

// Payment is the resolver for the payment field.
func (r *rideHistoryResolver) Payment(ctx context.Context, obj *db.RideHistory) (*db.PaymentHistory, error) {
	panic(fmt.Errorf("not implemented: Payment - payment"))
}

// Driver is the resolver for the driver field.
func (r *rideHistoryResolver) Driver(ctx context.Context, obj *db.RideHistory) (*db.Driver, error) {
	panic(fmt.Errorf("not implemented: Driver - driver"))
}

// User is the resolver for the user field.
func (r *rideHistoryResolver) User(ctx context.Context, obj *db.RideHistory) (*db.UserMetaData, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// DriverLocations is the resolver for the driverLocations field.
func (r *subscriptionResolver) DriverLocations(ctx context.Context, location *model.UserLocation) (<-chan []*model.CarLocation, error) {
    latLng := s2.LatLngFromDegrees(location.Lat,location.Lng);

    //Convert latlont to cellID
    tempCellID := s2.CellIDFromLatLng(latLng);

    //Get cellID at level 8;
    cellID := tempCellID.Parent(10);

    ch := make(chan []*model.CarLocation)
    fmt.Println("CellID")
    fmt.Println(tempCellID)
    fmt.Println(cellID);
    go func (){
        drivers := make(map[int64]*model.CarLocation)
        readerConfig:= kafkaconfig.NewKafkaReaderConfig("driver-location",[]string{r.Config.KafkaHost});

        reader:= kafka.NewReader(readerConfig)
        defer reader.Close()

        for{
            message, err := reader.ReadMessage(context.Background())
            if err != nil {
                log.Println("Error reading message:", err)
                break
            }


            var  loc *model.CarLocation
            e:=json.Unmarshal(message.Value,&loc)
            
            if e!=nil{
                fmt.Println(e);
                return;
            }

            drivers[loc.Driver.ID] = loc;
            fmt.Println(loc.Location.Long)
            list_drivers := []*model.CarLocation{}
            for _,v:=range drivers{
                list_drivers = append(list_drivers,v)
            }
            select{
            case ch<- list_drivers:
                fmt.Println("Send")
            }

            close(ch)
        }
    }();


    return ch,nil;
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// RideHistory returns graph.RideHistoryResolver implementation.
func (r *Resolver) RideHistory() graph.RideHistoryResolver { return &rideHistoryResolver{r} }

// Subscription returns graph.SubscriptionResolver implementation.
func (r *Resolver) Subscription() graph.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type rideHistoryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *rideHistoryResolver) PaymentID(ctx context.Context, obj *db.RideHistory) (*db.PaymentHistory, error) {
	panic(fmt.Errorf("not implemented: PaymentID - payment_id"))
}
func (r *userMetaDataResolver) RideHistory(ctx context.Context, obj *db.UserMetaData) ([]*db.RideHistory, error) {
	panic(fmt.Errorf("not implemented: RideHistory - rideHistory"))
}

type userMetaDataResolver struct{ *Resolver }
