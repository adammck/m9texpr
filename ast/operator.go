package ast

// TMP: This is just a placeholder until expressions are interfaces.
type Expr interface {
  String()
}

type Operator interface {
  Compare(left Expr, right Expr) bool
}

type Equals      struct { }
type NotEquals   struct { }
type GreaterThan struct { }
type LessThan    struct { }

// ==

func (o *Equals) Compare(left Expr, right Expr) bool {
  return false
}

func (o *Equals) String() string {
  return "op(eq)"
}

// !=

func (o *NotEquals) Compare(left Expr, right Expr) bool {
  return false
}

func (o *NotEquals) String() string {
  return "op(ne)"
}

// >

func (o *GreaterThan) Compare(left Expr, right Expr) bool {
  return false
}

func (o *GreaterThan) String() string {
  return "op(gt)"
}

// <

func (o *LessThan) Compare(left Expr, right Expr) bool {
  return false
}

func (o *LessThan) String() string {
  return "op(lt)"
}
