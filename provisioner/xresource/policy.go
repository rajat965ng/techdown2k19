package xresource

import (
	"../client"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"../structs"
	"strings"
)


func ResourceItem() *schema.Resource  {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
		 	"name": {
		 		Type: schema.TypeString,
		 		Required: true,
		 		ForceNew: true,
		 		Description: "Name of the policy",
			},
		 	"type": {
		 		Type: schema.TypeString,
		 		Required: true,
		 		Description: "Can be one of SecurityPolicy or License",
			},
		 	"description": {
		 		Type: schema.TypeString,
		 		Required: true,
		 		Description: "Description of the resource",
			},
		 	"rules": {
		 		Type:schema.TypeList,
		 		Required:true,
		 		Description: "Description of rules",
			},
		},
		Create:resourceCreateItem,
		Read:resourceReadItem,
		Update:resourceUpdateItem,
		Delete:resourceDeleteItem,
		Exists:resourceExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

var resourcePath = "api/v1/policies"




func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	apiClient:=(m).(*client.Client)

	log.Println("Calling resourceCreateItem")
	
	log.Println("name:",d.Get("name"))
	log.Println("type",d.Get("type"))
	log.Println("description",d.Get("description"))

	policy:=structs.Policy{

		Name: d.Get("name").(string),
		Type:d.Get("type").(string),
		Description:d.Get("description").(string),
	}
    err:=apiClient.NewItem(resourcePath,policy)

    if err!=nil{
    	return err
	}
    d.SetId(policy.Name)
	return nil
}

func resourceReadItem(d *schema.ResourceData, m interface{}) error {
	log.Println("Calling resourceReadItem")

	apiClient := m.(*client.Client)

	policyId := d.Id()
	log.Println("Invoke Api GetItem")
	policy, err := apiClient.GetItem(resourcePath,policyId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return fmt.Errorf("error finding Item with ID %s", policyId)
		}
	}

	d.SetId(policy.Name)
	d.Set("name", policy.Name)
	d.Set("description", policy.Description)
	d.Set("type", policy.Type)

	return nil
}

func resourceUpdateItem(d *schema.ResourceData, m interface{}) error {
	log.Println("Calling resourceUpdateItem")
	return nil
}

func resourceDeleteItem(d *schema.ResourceData, m interface{}) error {
	log.Println("Calling resourceDeleteItem")
	return nil
}

func resourceExistsItem(d *schema.ResourceData, m interface{}) (bool, error) {
	log.Println("Calling resourceExistsItem")
	return true, nil
}