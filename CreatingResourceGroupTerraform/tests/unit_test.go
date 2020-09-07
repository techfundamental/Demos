package tests

import (
	"../tests/models"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestUnit_Validate(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	terraform.Init(t, terraformOptions)
	validatedJson := terraform.RunTerraformCommand(t, terraformOptions, terraform.FormatArgs(terraformOptions, "validate", "-json")...)
	res := models.Validation{}

	json.Unmarshal([]byte(validatedJson), &res)
	assert.Equal(t, res.IsValid, true)
	assert.Equal(t, res.WarningCount, 0)
	assert.Equal(t, res.ErrorCount, 0)
}

func TestUnit_Plan(t *testing.T) {
	mydir, _ := os.Getwd()
	vf := path.Join(mydir, "vars/input.tfvars")
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		VarFiles: []string{vf},
	}

	tfPlanOutput := "terraform.tfplan"

	terraform.Init(t, terraformOptions)
	terraform.RunTerraformCommand(t, terraformOptions, terraform.FormatArgs(terraformOptions, "plan", "-out="+tfPlanOutput)...)
	terraformOptionsEmpty := &terraform.Options{}

	planJSON, err := terraform.RunTerraformCommandAndGetStdoutE(
		t, terraformOptions, terraform.FormatArgs(terraformOptionsEmpty, "show", "-json", tfPlanOutput)...,
	)

	if err != nil {
		t.Fatal(err)
	}

	res := models.Plan{}
	json.Unmarshal([]byte(planJSON), &res)
	assert.Equal(t, res.Planned.Outputs.Name.Value, "TerraformDemoTest-rg")
	assert.Equal(t, len(res.Planned.Modules.Deployments), 1)
	assert.Equal(t, res.Planned.Modules.Deployments[0].Type, "azurerm_resource_group")
}
