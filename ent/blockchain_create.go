// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shopspring/decimal"
	"github.com/skyisboss/pay-system/ent/blockchain"
	"github.com/skyisboss/pay-system/ent/schema"
)

// BlockchainCreate is the builder for creating a Blockchain entity.
type BlockchainCreate struct {
	config
	mutation *BlockchainMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BlockchainCreate) SetCreatedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableCreatedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BlockchainCreate) SetUpdatedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableUpdatedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetDeletedAt sets the "deleted_at" field.
func (bc *BlockchainCreate) SetDeletedAt(t time.Time) *BlockchainCreate {
	bc.mutation.SetDeletedAt(t)
	return bc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableDeletedAt(t *time.Time) *BlockchainCreate {
	if t != nil {
		bc.SetDeletedAt(*t)
	}
	return bc
}

// SetChain sets the "chain" field.
func (bc *BlockchainCreate) SetChain(s string) *BlockchainCreate {
	bc.mutation.SetChain(s)
	return bc
}

// SetTypes sets the "types" field.
func (bc *BlockchainCreate) SetTypes(s string) *BlockchainCreate {
	bc.mutation.SetTypes(s)
	return bc
}

// SetSymbol sets the "symbol" field.
func (bc *BlockchainCreate) SetSymbol(s string) *BlockchainCreate {
	bc.mutation.SetSymbol(s)
	return bc
}

// SetDecimals sets the "decimals" field.
func (bc *BlockchainCreate) SetDecimals(i int64) *BlockchainCreate {
	bc.mutation.SetDecimals(i)
	return bc
}

// SetStatus sets the "status" field.
func (bc *BlockchainCreate) SetStatus(i int64) *BlockchainCreate {
	bc.mutation.SetStatus(i)
	return bc
}

// SetTokenAddress sets the "token_address" field.
func (bc *BlockchainCreate) SetTokenAddress(s string) *BlockchainCreate {
	bc.mutation.SetTokenAddress(s)
	return bc
}

// SetTokenAbi sets the "token_abi" field.
func (bc *BlockchainCreate) SetTokenAbi(s string) *BlockchainCreate {
	bc.mutation.SetTokenAbi(s)
	return bc
}

// SetNillableTokenAbi sets the "token_abi" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableTokenAbi(s *string) *BlockchainCreate {
	if s != nil {
		bc.SetTokenAbi(*s)
	}
	return bc
}

// SetColdAddress sets the "cold_address" field.
func (bc *BlockchainCreate) SetColdAddress(s string) *BlockchainCreate {
	bc.mutation.SetColdAddress(s)
	return bc
}

// SetHotAddress sets the "hot_address" field.
func (bc *BlockchainCreate) SetHotAddress(s string) *BlockchainCreate {
	bc.mutation.SetHotAddress(s)
	return bc
}

// SetScanBlockNum sets the "scan_block_num" field.
func (bc *BlockchainCreate) SetScanBlockNum(i int64) *BlockchainCreate {
	bc.mutation.SetScanBlockNum(i)
	return bc
}

// SetMinFreeNum sets the "min_free_num" field.
func (bc *BlockchainCreate) SetMinFreeNum(i int64) *BlockchainCreate {
	bc.mutation.SetMinFreeNum(i)
	return bc
}

// SetMinConfirmNum sets the "min_confirm_num" field.
func (bc *BlockchainCreate) SetMinConfirmNum(i int64) *BlockchainCreate {
	bc.mutation.SetMinConfirmNum(i)
	return bc
}

// SetWithdrawFee sets the "withdraw_fee" field.
func (bc *BlockchainCreate) SetWithdrawFee(f float64) *BlockchainCreate {
	bc.mutation.SetWithdrawFee(f)
	return bc
}

// SetWithdrawFeeType sets the "withdraw_fee_type" field.
func (bc *BlockchainCreate) SetWithdrawFeeType(i int64) *BlockchainCreate {
	bc.mutation.SetWithdrawFeeType(i)
	return bc
}

// SetMinDeposit sets the "min_deposit" field.
func (bc *BlockchainCreate) SetMinDeposit(d decimal.Decimal) *BlockchainCreate {
	bc.mutation.SetMinDeposit(d)
	return bc
}

// SetMinWithdraw sets the "min_withdraw" field.
func (bc *BlockchainCreate) SetMinWithdraw(d decimal.Decimal) *BlockchainCreate {
	bc.mutation.SetMinWithdraw(d)
	return bc
}

// SetMinCollect sets the "min_collect" field.
func (bc *BlockchainCreate) SetMinCollect(d decimal.Decimal) *BlockchainCreate {
	bc.mutation.SetMinCollect(d)
	return bc
}

// SetGasPrice sets the "gas_price" field.
func (bc *BlockchainCreate) SetGasPrice(sp schema.GasPrice) *BlockchainCreate {
	bc.mutation.SetGasPrice(sp)
	return bc
}

// SetNillableGasPrice sets the "gas_price" field if the given value is not nil.
func (bc *BlockchainCreate) SetNillableGasPrice(sp *schema.GasPrice) *BlockchainCreate {
	if sp != nil {
		bc.SetGasPrice(*sp)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BlockchainCreate) SetID(u uint64) *BlockchainCreate {
	bc.mutation.SetID(u)
	return bc
}

// Mutation returns the BlockchainMutation object of the builder.
func (bc *BlockchainCreate) Mutation() *BlockchainMutation {
	return bc.mutation
}

// Save creates the Blockchain in the database.
func (bc *BlockchainCreate) Save(ctx context.Context) (*Blockchain, error) {
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockchainCreate) SaveX(ctx context.Context) *Blockchain {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockchainCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockchainCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockchainCreate) check() error {
	if _, ok := bc.mutation.Chain(); !ok {
		return &ValidationError{Name: "chain", err: errors.New(`ent: missing required field "Blockchain.chain"`)}
	}
	if _, ok := bc.mutation.Types(); !ok {
		return &ValidationError{Name: "types", err: errors.New(`ent: missing required field "Blockchain.types"`)}
	}
	if _, ok := bc.mutation.Symbol(); !ok {
		return &ValidationError{Name: "symbol", err: errors.New(`ent: missing required field "Blockchain.symbol"`)}
	}
	if _, ok := bc.mutation.Decimals(); !ok {
		return &ValidationError{Name: "decimals", err: errors.New(`ent: missing required field "Blockchain.decimals"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Blockchain.status"`)}
	}
	if _, ok := bc.mutation.TokenAddress(); !ok {
		return &ValidationError{Name: "token_address", err: errors.New(`ent: missing required field "Blockchain.token_address"`)}
	}
	if _, ok := bc.mutation.ColdAddress(); !ok {
		return &ValidationError{Name: "cold_address", err: errors.New(`ent: missing required field "Blockchain.cold_address"`)}
	}
	if _, ok := bc.mutation.HotAddress(); !ok {
		return &ValidationError{Name: "hot_address", err: errors.New(`ent: missing required field "Blockchain.hot_address"`)}
	}
	if _, ok := bc.mutation.ScanBlockNum(); !ok {
		return &ValidationError{Name: "scan_block_num", err: errors.New(`ent: missing required field "Blockchain.scan_block_num"`)}
	}
	if _, ok := bc.mutation.MinFreeNum(); !ok {
		return &ValidationError{Name: "min_free_num", err: errors.New(`ent: missing required field "Blockchain.min_free_num"`)}
	}
	if _, ok := bc.mutation.MinConfirmNum(); !ok {
		return &ValidationError{Name: "min_confirm_num", err: errors.New(`ent: missing required field "Blockchain.min_confirm_num"`)}
	}
	if _, ok := bc.mutation.WithdrawFee(); !ok {
		return &ValidationError{Name: "withdraw_fee", err: errors.New(`ent: missing required field "Blockchain.withdraw_fee"`)}
	}
	if _, ok := bc.mutation.WithdrawFeeType(); !ok {
		return &ValidationError{Name: "withdraw_fee_type", err: errors.New(`ent: missing required field "Blockchain.withdraw_fee_type"`)}
	}
	if _, ok := bc.mutation.MinDeposit(); !ok {
		return &ValidationError{Name: "min_deposit", err: errors.New(`ent: missing required field "Blockchain.min_deposit"`)}
	}
	if _, ok := bc.mutation.MinWithdraw(); !ok {
		return &ValidationError{Name: "min_withdraw", err: errors.New(`ent: missing required field "Blockchain.min_withdraw"`)}
	}
	if _, ok := bc.mutation.MinCollect(); !ok {
		return &ValidationError{Name: "min_collect", err: errors.New(`ent: missing required field "Blockchain.min_collect"`)}
	}
	return nil
}

func (bc *BlockchainCreate) sqlSave(ctx context.Context) (*Blockchain, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlockchainCreate) createSpec() (*Blockchain, *sqlgraph.CreateSpec) {
	var (
		_node = &Blockchain{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blockchain.Table, sqlgraph.NewFieldSpec(blockchain.FieldID, field.TypeUint64))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(blockchain.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(blockchain.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.DeletedAt(); ok {
		_spec.SetField(blockchain.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := bc.mutation.Chain(); ok {
		_spec.SetField(blockchain.FieldChain, field.TypeString, value)
		_node.Chain = value
	}
	if value, ok := bc.mutation.Types(); ok {
		_spec.SetField(blockchain.FieldTypes, field.TypeString, value)
		_node.Types = value
	}
	if value, ok := bc.mutation.Symbol(); ok {
		_spec.SetField(blockchain.FieldSymbol, field.TypeString, value)
		_node.Symbol = value
	}
	if value, ok := bc.mutation.Decimals(); ok {
		_spec.SetField(blockchain.FieldDecimals, field.TypeInt64, value)
		_node.Decimals = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(blockchain.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.TokenAddress(); ok {
		_spec.SetField(blockchain.FieldTokenAddress, field.TypeString, value)
		_node.TokenAddress = value
	}
	if value, ok := bc.mutation.TokenAbi(); ok {
		_spec.SetField(blockchain.FieldTokenAbi, field.TypeString, value)
		_node.TokenAbi = value
	}
	if value, ok := bc.mutation.ColdAddress(); ok {
		_spec.SetField(blockchain.FieldColdAddress, field.TypeString, value)
		_node.ColdAddress = value
	}
	if value, ok := bc.mutation.HotAddress(); ok {
		_spec.SetField(blockchain.FieldHotAddress, field.TypeString, value)
		_node.HotAddress = value
	}
	if value, ok := bc.mutation.ScanBlockNum(); ok {
		_spec.SetField(blockchain.FieldScanBlockNum, field.TypeInt64, value)
		_node.ScanBlockNum = value
	}
	if value, ok := bc.mutation.MinFreeNum(); ok {
		_spec.SetField(blockchain.FieldMinFreeNum, field.TypeInt64, value)
		_node.MinFreeNum = value
	}
	if value, ok := bc.mutation.MinConfirmNum(); ok {
		_spec.SetField(blockchain.FieldMinConfirmNum, field.TypeInt64, value)
		_node.MinConfirmNum = value
	}
	if value, ok := bc.mutation.WithdrawFee(); ok {
		_spec.SetField(blockchain.FieldWithdrawFee, field.TypeFloat64, value)
		_node.WithdrawFee = value
	}
	if value, ok := bc.mutation.WithdrawFeeType(); ok {
		_spec.SetField(blockchain.FieldWithdrawFeeType, field.TypeInt64, value)
		_node.WithdrawFeeType = value
	}
	if value, ok := bc.mutation.MinDeposit(); ok {
		_spec.SetField(blockchain.FieldMinDeposit, field.TypeFloat64, value)
		_node.MinDeposit = value
	}
	if value, ok := bc.mutation.MinWithdraw(); ok {
		_spec.SetField(blockchain.FieldMinWithdraw, field.TypeFloat64, value)
		_node.MinWithdraw = value
	}
	if value, ok := bc.mutation.MinCollect(); ok {
		_spec.SetField(blockchain.FieldMinCollect, field.TypeFloat64, value)
		_node.MinCollect = value
	}
	if value, ok := bc.mutation.GasPrice(); ok {
		_spec.SetField(blockchain.FieldGasPrice, field.TypeJSON, value)
		_node.GasPrice = value
	}
	return _node, _spec
}

// BlockchainCreateBulk is the builder for creating many Blockchain entities in bulk.
type BlockchainCreateBulk struct {
	config
	err      error
	builders []*BlockchainCreate
}

// Save creates the Blockchain entities in the database.
func (bcb *BlockchainCreateBulk) Save(ctx context.Context) ([]*Blockchain, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blockchain, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockchainMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockchainCreateBulk) SaveX(ctx context.Context) []*Blockchain {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockchainCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockchainCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}