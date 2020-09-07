package tests

import (
	"os"
	"path"
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	mydir, _ := os.Getwd()
	vf := path.Join(mydir, "vars/input.tfvars")
	terraformOptions := &terraform.Options{
		TerraformDir: "../",
		VarFiles: []string{vf},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	componentName := terraform.Output(t, terraformOptions, "name")
	assert.Equal(t, componentName, "TerraformDemoTest-rg")
}
