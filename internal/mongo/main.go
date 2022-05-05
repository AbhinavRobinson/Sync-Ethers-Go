package mongo

import (
	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	config "sync-ethers-go/config"
	internal "sync-ethers-go/internal"
)

func InitMongo() {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "sync-ethers", options.Client().ApplyURI(config.MONGO_URL))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize mongo")
	}
	log.Debug().Msg("Mongo initialized")
}

func NewContract(address string, contractType string) *internal.StoredContracts {
	return &internal.StoredContracts{
		Address: address,
		Type:    contractType,
	}
}

func GetContractsFromDB() []internal.StoredContracts {
	// get all from db
	results := []internal.StoredContracts{}
	col := mgm.Coll(&internal.StoredContracts{})
	cursor, err := col.Find(mgm.Ctx(), bson.M{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get contracts from DB")
	}
	for cursor.Next(mgm.Ctx()) {
		var contract internal.StoredContracts
		err := cursor.Decode(&contract)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to decode contract from DB")
		}
		results = append(results, contract)
	}
	return results
}

func AddContractToDB(contract *internal.StoredContracts) bool {
	// check if exists in db
	c := &internal.StoredContracts{}
	err := mgm.Coll(c).FindOne(mgm.Ctx(), bson.M{"address": contract.Address}).Decode(c)
	if err != nil {
		log.Debug().Msgf("Contract %s not found in db", contract.Address)
		// insert
		err := mgm.Coll(contract).Create(contract)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to add contract to DB")
			return false
		}
		return true
	}
	log.Debug().Msgf("Contract %s found in db", contract.Address)
	return true
}

func DeleteContractFromDB(contract *internal.StoredContracts) bool {
	// delete from db
	res, err := mgm.Coll(contract).DeleteOne(mgm.Ctx(), bson.M{"address": contract.Address})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to delete contract from DB")
		return false
	}
	if res.DeletedCount == 0 {
		log.Debug().Msgf("Contract %s not found in db", contract.Address)
		return false
	}
	return true
}
