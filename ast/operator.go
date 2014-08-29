package ast

type Operator interface {
  Compare(left Expr, right Expr) (Bool, error)
}

type Equals      struct { }
type NotEquals   struct { }
type GreaterThan struct { }
type LessThan    struct { }

// ==

func (o *Equals) Compare(left Expr, right Expr) (Bool, error) {
  return MakeBool(left == right), nil
}

func (o *Equals) String() string {
  return "op(eq)"
}

// !=

func (o *NotEquals) Compare(left Expr, right Expr) (Bool, error) {
  return MakeBool(left != right), nil
}

func (o *NotEquals) String() string {
  return "op(ne)"
}

// >

func (o *GreaterThan) Compare(left Expr, right Expr) (Bool, error) {
  return MakeBool(false), nil
}

func (o *GreaterThan) String() string {
  return "op(gt)"
}

// <

func (o *LessThan) Compare(left Expr, right Expr) (Bool, error) {
  return MakeBool(false), nil
}

func (o *LessThan) String() string {
  return "op(lt)"
}
