package provider

import (
	"github.com/bigcommerce/terraform-provider-cortex/internal/cortex"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

/***********************************************************************************************************************
 * Models
 **********************************************************************************************************************/

// DepartmentResourceModel describes the department data model within Terraform.
type DepartmentResourceModel struct {
	Id          types.String                    `tfsdk:"id"`
	Tag         types.String                    `tfsdk:"tag"`
	Name        types.String                    `tfsdk:"name"`
	Description types.String                    `tfsdk:"description"`
	Members     []DepartmentMemberResourceModel `tfsdk:"members"`
}

func (r *DepartmentResourceModel) FromApiModel(entity cortex.Department) {
	r.Id = types.StringValue(entity.Tag)
	r.Tag = types.StringValue(entity.Tag)
	r.Name = types.StringValue(entity.Name)
	r.Description = types.StringValue(entity.Description)
	if entity.Members != nil {
		r.Members = make([]DepartmentMemberResourceModel, len(entity.Members))
		for i, member := range entity.Members {
			m := DepartmentMemberResourceModel{}
			r.Members[i] = m.FromApiModel(&member)
		}
	}
}

func (r *DepartmentResourceModel) ToApiModel() cortex.Department {
	entity := cortex.Department{
		Tag:         r.Tag.ValueString(),
		Name:        r.Name.ValueString(),
		Description: r.Description.ValueString(),
	}
	var members []cortex.DepartmentMember
	for _, member := range r.Members {
		members = append(members, member.ToApiModel())
	}
	entity.Members = members
	return entity
}

// ToUpdateRequest https://docs.cortex.io/docs/api/update-department
func (r *DepartmentResourceModel) ToUpdateRequest() cortex.UpdateDepartmentRequest {
	var members []cortex.DepartmentMember
	for _, member := range r.Members {
		members = append(members, member.ToApiModel())
	}
	return cortex.UpdateDepartmentRequest{
		Name:        r.Name.ValueString(),
		Description: r.Description.ValueString(),
		Members:     members,
	}
}

/***********************************************************************************************************************
 * Members
 **********************************************************************************************************************/

type DepartmentMemberResourceModel struct {
	Name        types.String `tfsdk:"name"`
	Email       types.String `tfsdk:"email"`
	Description types.String `tfsdk:"description"`
}

func (o *DepartmentMemberResourceModel) ToApiModel() cortex.DepartmentMember {
	return cortex.DepartmentMember{
		Name:        o.Name.ValueString(),
		Email:       o.Email.ValueString(),
		Description: o.Description.ValueString(),
	}
}

func (o *DepartmentMemberResourceModel) FromApiModel(member *cortex.DepartmentMember) DepartmentMemberResourceModel {
	return DepartmentMemberResourceModel{
		Name:        types.StringValue(member.Name),
		Email:       types.StringValue(member.Email),
		Description: types.StringValue(member.Description),
	}
}

/***********************************************************************************************************************
 * Data Source
 **********************************************************************************************************************/

// DepartmentDataSourceModel describes the data source data model.
type DepartmentDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Tag         types.String `tfsdk:"tag"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func (o *DepartmentDataSourceModel) FromApiModel(entity cortex.Department) {
	o.Id = types.StringValue(entity.Tag)
	o.Tag = types.StringValue(entity.Tag)
	o.Name = types.StringValue(entity.Name)
	o.Description = types.StringValue(entity.Description)
}
