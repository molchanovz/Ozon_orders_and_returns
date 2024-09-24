package main

import (
	"encoding/json"
	"log"
)

func get_returns_fbo(Client_Id, Api_Key string, LastID int) ([]ReturnFBO, int) {
	var returns ReturnsFBO
	jsonString := v3_returns_company_fbo(Client_Id, Api_Key, LastID) // assuming this function returns the JSON string
	err := json.Unmarshal([]byte(jsonString), &returns)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return returns.Returns, returns.LastID
}
func get_posting_fbo(Client_Id, Api_Key, PostingNumber string) PostingFBO {
	var posting PostingFBO
	jsonString := v2_posting_fbo_get(Client_Id, Api_Key, PostingNumber) // assuming this function returns the JSON string
	err := json.Unmarshal([]byte(jsonString), &posting)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return posting
}

func get_returns_fbs(Client_Id, Api_Key string, LastID int) ([]ReturnFBS, int) {
	var returns ReturnsFBS
	jsonString := v3_returns_company_fbs(Client_Id, Api_Key, LastID) // assuming this function returns the JSON string
	err := json.Unmarshal([]byte(jsonString), &returns)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return returns.Returns, returns.LastID
}
func get_posting_fbs(Client_Id, Api_Key, PostingNumber string) PostingFBS {
	var posting PostingFBS
	jsonString := v3_posting_fbs_get(Client_Id, Api_Key, PostingNumber) // assuming this function returns the JSON string
	err := json.Unmarshal([]byte(jsonString), &posting)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return posting
}

func get_postings_list_fbs(Client_Id, Api_Key string) PostingsList_FBS {
	var postingList PostingsList_FBS
	jsonString := v3_posting_fbs_list(Client_Id, Api_Key)
	err := json.Unmarshal([]byte(jsonString), &postingList)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return postingList
}
func get_postings_list_fbo(Client_Id, Api_Key string) PostingsList_FBO {
	var postingList PostingsList_FBO
	jsonString := v2_posting_fbo_list(Client_Id, Api_Key)
	err := json.Unmarshal([]byte(jsonString), &postingList)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return postingList
}
