package xprovider

import (
	"../client"
	"../xresource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func  Provider() terraform.ResourceProvider  {
	log.Print("Invoking provider ")
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host":{
				Type:schema.TypeString,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("SERVICE_HOST",""),
			},
			"port":{
				Type:schema.TypeInt,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("SERVICE_PORT",""),
			},
			"username":{
				Type:schema.TypeString,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("SERVICE_USERNAME",""),
			},
			"password":{
				Type:schema.TypeString,
				Required:true,
				DefaultFunc:schema.EnvDefaultFunc("SERVICE_PASSWD",""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"xray_policy": xresource.ResourceItem(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{},error)  {
	log.Println("Invoking providerConfigure")
	host:=data.Get("host").(string)
	port:=data.Get("port").(int)
	username:=data.Get("username").(string)
	password:=data.Get("password").(string)
	return client.NewClient(host,port,username,password,data.Id()), nil
}