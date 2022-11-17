package testschema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema/fwxschema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ fwxschema.BlockWithSetValidators = BlockWithSetValidators{}

type BlockWithSetValidators struct {
	Attributes          map[string]fwschema.Attribute
	Blocks              map[string]fwschema.Block
	DeprecationMessage  string
	Description         string
	MarkdownDescription string
	MaxItems            int64
	MinItems            int64
	Validators          []validator.Set
}

// ApplyTerraform5AttributePathStep satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (any, error) {
	return b.Type().ApplyTerraform5AttributePathStep(step)
}

// Equal satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) Equal(o fwschema.Block) bool {
	_, ok := o.(BlockWithSetValidators)

	if !ok {
		return false
	}

	return fwschema.BlocksEqual(b, o)
}

// GetAttributes satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetAttributes() map[string]fwschema.Attribute {
	return nil
}

// GetBlocks satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetBlocks() map[string]fwschema.Block {
	return nil
}

// GetDeprecationMessage satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetDeprecationMessage() string {
	return b.DeprecationMessage
}

// GetDescription satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetDescription() string {
	return b.Description
}

// GetMarkdownDescription satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetMarkdownDescription() string {
	return b.MarkdownDescription
}

// GetMaxItems satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetMaxItems() int64 {
	return b.MaxItems
}

// GetMinItems satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetMinItems() int64 {
	return b.MinItems
}

// GetNestingMode satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) GetNestingMode() fwschema.BlockNestingMode {
	return fwschema.BlockNestingModeSet
}

// SetValidators satisfies the fwxschema.BlockWithSetValidators interface.
func (b BlockWithSetValidators) SetValidators() []validator.Set {
	return b.Validators
}

// Type satisfies the fwschema.Block interface.
func (b BlockWithSetValidators) Type() attr.Type {
	attrType := types.ObjectType{
		AttrTypes: make(map[string]attr.Type, len(b.GetAttributes())+len(b.GetBlocks())),
	}

	for attrName, attr := range b.GetAttributes() {
		attrType.AttrTypes[attrName] = attr.FrameworkType()
	}

	for blockName, block := range b.GetBlocks() {
		attrType.AttrTypes[blockName] = block.Type()
	}

	return types.SetType{
		ElemType: attrType,
	}
}