package bug

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/property"
	"entgo.io/bug/ent/user"
)

func (r *contractResolver) Tenant(ctx context.Context, obj *ent.Contract) (*ent.User, error) {
	propertyContract, err := obj.Property(ctx)

	if err != nil {
		log.Printf("failed querying contract property: %v", err)
		return nil, err
	}

	return obj.
		QueryUsers().
		Where(user.Not(user.HasPropertiesWith(property.ID(propertyContract.ID)))).
		Only(ctx)
}

func (r *mutationResolver) CreateUser(ctx context.Context, user UserInput) (*ent.User, error) {
	newUser, err := r.client.User.Create().
		SetNames(user.Names).
		SetLastnames(user.Lastnames).
		SetEmail(user.Email).
		SetActivate(true).
		SetCreatedAt(time.Now()).
		SetBirthday(*user.Birthday).Save(ctx)

	if err != nil {
		log.Printf("failed creating user: %v", err)
		return nil, err
	}

	return newUser, err
}

func (r *mutationResolver) CreateProperty(ctx context.Context, property PropertyInput) (*ent.Property, error) {
	newProperty, err := r.client.Property.Create().
		SetClass(property.Class).
		SetDeleted(false).
		SetAddress(property.Address).
		SetCity(property.City).
		SetDescription(property.Description).
		SetOwnerID(property.OwnerID).Save(ctx)

	if err != nil {
		log.Printf("failed creating property: %v", err)
		return nil, err
	}

	return newProperty, err
}

func (r *mutationResolver) CreateContract(ctx context.Context, contract ContractInput) (*ent.Contract, error) {
	newContract, err := r.client.Contract.Create().
		SetStartDate(contract.StartDate).
		SetEndDate(contract.EndDate).
		SetPayAmount(contract.PayAmount).
		SetPayDate(contract.PayDate).
		SetPropertyID(contract.PropertyID).Save(ctx)

	if err != nil {
		log.Printf("failed creating contract: %v", err)
		return nil, err
	}

	err = r.client.User.UpdateOneID(contract.OwnerID).AddContractIDs(newContract.ID).Exec(ctx)
	if err != nil {
		log.Printf("failed updating contract owner: %v", err)
		return nil, err
	}

	err = r.client.User.UpdateOneID(contract.TenantID).AddContractIDs(newContract.ID).Exec(ctx)
	if err != nil {
		log.Printf("failed updating contract tenant: %v", err)
		return nil, err
	}

	return newContract, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*ent.User, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		log.Printf("failed get all users: %v", err)
		return nil, err
	}

	return users, err
}

func (r *queryResolver) GetProperty(ctx context.Context, id int) (*ent.Property, error) {
	property, err := r.client.Property.Get(ctx, id)
	if err != nil {
		log.Printf("failed get property: %v", err)
		return nil, err
	}

	return property, err
}

func (r *queryResolver) GetContract(ctx context.Context, id int) (*ent.Contract, error) {
	property, err := r.client.Contract.Get(ctx, id)
	if err != nil {
		log.Printf("failed get contract: %v", err)
		return nil, err
	}

	return property, err
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Contract returns ContractResolver implementation.
func (r *Resolver) Contract() ContractResolver { return &contractResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type contractResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
