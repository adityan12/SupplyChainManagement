package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SupplyChainContract struct {
}

type User struct {
	UserID    string `json:"UserID"`
	Name      string `json:"Name"`
	Email     string `json:"Email"`
	Phone     string `json:"Phone"`
	Address   string `json:"Address"`
	UserType  string `json:"UserType"`
	AuthCode  string `json:"AuthCode"`
}

type TrackingEvent struct {
	Timestamp time.Time `json:"Timestamp"`
	Event     string    `json:"Event"`
}

type Product struct {
	ProductID   string         `json:"ProductID"`
	Name        string         `json:"Name"`
	Description string         `json:"Description"`
	Manufacturer string         `json:"Manufacturer"`
	SenderName  string         `json:"SenderName"`
	AuthCode    string         `json:"AuthCode"`
	Tracking    []TrackingEvent `json:"Tracking"`
}

func main() {
	err := shim.Start(new(SupplyChainContract))
	if err != nil {
		fmt.Printf("Error starting supply chain contract: %s", err)
	}
}

func (t *SupplyChainContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SupplyChainContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "createUser" {
		return t.createUser(stub, args)
	} else if function == "createProduct" {
		return t.createProduct(stub, args)
	} else if function == "addTrackingEvent" {
		return t.addTrackingEvent(stub, args)
	} else if function == "queryProduct" {
		return t.queryProduct(stub, args)
	}

	return shim.Error("Invalid function name")
}

func (t *SupplyChainContract) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expected 7")
	}

	userID := args[0]
	name := args[1]
	email := args[2]
	phone := args[3]
	address := args[4]
	userType := args[5]
	authCode := args[6]

	user := User{
		UserID:    userID,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Address:   address,
		UserType:  userType,
		AuthCode:  authCode,
	}

	userAsBytes, _ := json.Marshal(user)
	err := stub.PutState(userID, userAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SupplyChainContract) createProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expected 6")
	}

	productID := args[0]
	name := args[1]
	description := args[2]
	manufacturer := args[3]
	senderName := args[4]
	authCode := args[5]

	product := Product{
		ProductID:   productID,
		Name:        name,
		Description: description,
		Manufacturer: manufacturer,
		SenderName:  senderName,
		AuthCode:    authCode,
	}

	productAsBytes, _ := json.Marshal(product)
	err := stub.PutState(productID, productAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SupplyChainContract) addTrackingEvent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expected 2")
	}

	productID := args[0]
	event := args[1]

	productAsBytes, err := stub.GetState(productID)
	if err != nil {
		return shim.Error(err.Error())
	}

	if productAsBytes == nil {
		return shim.Error("Product not found")
	}

	var product Product
	err = json.Unmarshal(productAsBytes, &product)
	if err != nil {
		return shim.Error(err.Error())
	}

	trackingEvent := TrackingEvent{
		Timestamp: time.Now(),
		Event:     event,
	}

	product.Tracking = append(product.Tracking, trackingEvent)

	productAsBytes, err = json.Marshal(product)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(productID, productAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SupplyChainContract) queryProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expected 1")
	}

	productID := args[0]

	productAsBytes, err := stub.GetState(productID)
	if err != nil {
		return shim.Error(err.Error())
	}

	if productAsBytes == nil {
		return shim.Error("Product not found")
	}

	return shim.Success(productAsBytes)
}
