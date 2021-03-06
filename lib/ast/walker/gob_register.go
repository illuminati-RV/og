package walker

import (
	"encoding/gob"
	"github.com/champii/og/lib/ast"
)

func RunGobRegister() {
	gob.Register(&Template{})
	gob.Register(&TemplateSerie{})
	gob.Register(&ast.SourceFile{})
	gob.Register(&ast.Package{})
	gob.Register(&ast.Import{})
	gob.Register(&ast.ImportSpec{})
	gob.Register(&ast.TopLevel{})
	gob.Register(&ast.Declaration{})
	gob.Register(&ast.ConstDecl{})
	gob.Register(&ast.ConstSpec{})
	gob.Register(&ast.ExpressionList{})
	gob.Register(&ast.Parameters{})
	gob.Register(&ast.TypeDecl{})
	gob.Register(&ast.TypeSpec{})
	gob.Register(&ast.FunctionDecl{})
	gob.Register(&ast.Function{})
	gob.Register(&ast.MethodDecl{})
	gob.Register(&ast.Receiver{})
	gob.Register(&ast.VarDecl{})
	gob.Register(&ast.VarSpec{})
	gob.Register(&ast.Block{})
	gob.Register(&ast.Statement{})
	gob.Register(&ast.SimpleStmt{})
	gob.Register(&ast.SendStmt{})
	gob.Register(&ast.IncDecStmt{})
	gob.Register(&ast.Assignment{})
	gob.Register(&ast.ShortVarDecl{})
	gob.Register(&ast.LabeledStmt{})
	gob.Register(&ast.ReturnStmt{})
	gob.Register(&ast.BreakStmt{})
	gob.Register(&ast.ContinueStmt{})
	gob.Register(&ast.GotoStmt{})
	gob.Register(&ast.FallthroughStmt{})
	gob.Register(&ast.DeferStmt{})
	gob.Register(&ast.IfStmt{})
	gob.Register(&ast.SwitchStmt{})
	gob.Register(&ast.ExprSwitchStmt{})
	gob.Register(&ast.ExprCaseClause{})
	gob.Register(&ast.ExprSwitchCase{})
	gob.Register(&ast.TypeSwitchStmt{})
	gob.Register(&ast.TypeSwitchGuard{})
	gob.Register(&ast.TypeCaseClause{})
	gob.Register(&ast.TypeSwitchCase{})
	gob.Register(&ast.SelectStmt{})
	gob.Register(&ast.CommClause{})
	gob.Register(&ast.CommCase{})
	gob.Register(&ast.RecvStmt{})
	gob.Register(&ast.ForStmt{})
	gob.Register(&ast.ForClause{})
	gob.Register(&ast.RangeClause{})
	gob.Register(&ast.GoStmt{})
	gob.Register(&ast.Type{})
	gob.Register(&ast.TypeLit{})
	gob.Register(&ast.ArrayType{})
	gob.Register(&ast.PointerType{})
	gob.Register(&ast.InterfaceType{})
	gob.Register(&ast.SliceType{})
	gob.Register(&ast.MapType{})
	gob.Register(&ast.ChannelType{})
	gob.Register(&ast.MethodSpec{})
	gob.Register(&ast.FunctionType{})
	gob.Register(&ast.Signature{})
	gob.Register(&ast.TemplateSpec{})
	gob.Register(&ast.Result{})
	gob.Register(&ast.Parameter{})
	gob.Register(&ast.Operand{})
	gob.Register(&ast.Literal{})
	gob.Register(&ast.OperandName{})
	gob.Register(&ast.CompositeLit{})
	gob.Register(&ast.LiteralType{})
	gob.Register(&ast.LiteralValue{})
	gob.Register(&ast.KeyedElement{})
	gob.Register(&ast.Key{})
	gob.Register(&ast.Element{})
	gob.Register(&ast.StructType{})
	gob.Register(&ast.FieldDecl{})
	gob.Register(&ast.IdentifierList{})
	gob.Register(&ast.InlineStructMethod{})
	gob.Register(&ast.AnonymousField{})
	gob.Register(&ast.FunctionLit{})
	gob.Register(&ast.PrimaryExpr{})
	gob.Register(&ast.SecondaryExpr{})
	gob.Register(&ast.Index{})
	gob.Register(&ast.Slice{})
	gob.Register(&ast.TypeAssertion{})
	gob.Register(&ast.Arguments{})
	gob.Register(&ast.MethodExpr{})
	gob.Register(&ast.ReceiverType{})
	gob.Register(&ast.Expression{})
	gob.Register(&ast.UnaryExpr{})
	gob.Register(&ast.Conversion{})
	gob.Register(&ast.Interpret{})
}
