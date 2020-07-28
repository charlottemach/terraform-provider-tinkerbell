package tinkerbell

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/tinkerbell/tink/protos/hardware"
)

func resourceHardware() *schema.Resource {
	return &schema.Resource{
		Create: resourceHardwareCreate,
		Read:   resourceHardwareRead,
		Delete: resourceHardwareDelete,
		Schema: map[string]*schema.Schema{
			"data": {
				Type:     schema.TypeString,
				Required: true,
                ForceNew: true,
			},
		},
	}
}

func resourceHardwareCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*TinkClient).HardwareClient

	req := hardware.PushRequest{
        Data: d.Get("data").(string),
	}

	_ , err := c.Push(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("pushing hardware failed: %w", err)
	}

	return nil
}

func resourceHardwareRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*TinkClient).HardwareClient

	req := hardware.GetRequest{
        ID: d.Id(),
	}

	_, err := c.ByID(context.Background(), &req)
	if err != nil {
		return fmt.Errorf("getting hardware by ID failed: %w", err)
	}

	return nil
}

func resourceHardwareDelete(d *schema.ResourceData, m interface{}) error {
	//c := m.(*TinkClient).HardwareClient
    //
	//req := hardware.DeleteRequest{
	//	Id: d.Id(),
	//}

	//if _, err := c.Delete(context.Background(), &req); err != nil {
	//	return fmt.Errorf("removing hardware failed: %w", err)
	//}

	return nil
}
