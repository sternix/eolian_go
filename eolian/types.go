package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type FunctionType int

const (
	FunctionTypeUnresolved FunctionType = C.EOLIAN_UNRESOLVED
	FunctionTypeProperty   FunctionType = C.EOLIAN_PROPERTY
	FunctionTypePropSet    FunctionType = C.EOLIAN_PROP_SET
	FunctionTypePropGet    FunctionType = C.EOLIAN_PROP_GET
	FunctionTypeMethod     FunctionType = C.EOLIAN_METHOD
)

type ParameterDir int

const (
	// not in 1.18.4
	//ParameterDirUnknown ParameterDir = C.EOLIAN_UNKNOWN_PARAM
	ParameterDirIn    ParameterDir = C.EOLIAN_IN_PARAM
	ParameterDirOut   ParameterDir = C.EOLIAN_OUT_PARAM
	ParameterDirInOut ParameterDir = C.EOLIAN_INOUT_PARAM
)

type ClassType int

const (
	ClassTypeUnknown   ClassType = C.EOLIAN_CLASS_UNKNOWN_TYPE
	ClassTypeRegular   ClassType = C.EOLIAN_CLASS_REGULAR
	ClassTypeAbstract  ClassType = C.EOLIAN_CLASS_ABSTRACT
	ClassTypeMixin     ClassType = C.EOLIAN_CLASS_MIXIN
	ClassTypeInterface ClassType = C.EOLIAN_CLASS_INTERFACE
)

type ObjectScope int

const (
	// not in 1.18.4
	// ObjectScopeUnknown   ObjectScope = C.EOLIAN_SCOPE_UNKNOWN
	ObjectScopePublic    ObjectScope = C.EOLIAN_SCOPE_PUBLIC
	ObjectScopePrivate   ObjectScope = C.EOLIAN_SCOPE_PRIVATE
	ObjectScopeProtected ObjectScope = C.EOLIAN_SCOPE_PROTECTED
)

type TypedeclType int

const (
	TypedeclTypeUnknown      TypedeclType = C.EOLIAN_TYPEDECL_UNKNOWN
	TypedeclTypeStruct       TypedeclType = C.EOLIAN_TYPEDECL_STRUCT
	TypedeclTypeStructOpaque TypedeclType = C.EOLIAN_TYPEDECL_STRUCT_OPAQUE
	TypedeclTypeEnum         TypedeclType = C.EOLIAN_TYPEDECL_ENUM
	TypedeclTypeAlias        TypedeclType = C.EOLIAN_TYPEDECL_ALIAS
)

type EoType int

const (
	EoTypeUnknown         EoType = C.EOLIAN_TYPE_UNKNOWN_TYPE
	EoTypeVoid            EoType = C.EOLIAN_TYPE_VOID
	EoTypeRegular         EoType = C.EOLIAN_TYPE_REGULAR
	EoTypeComplex         EoType = C.EOLIAN_TYPE_COMPLEX
	EoTypeClass           EoType = C.EOLIAN_TYPE_CLASS
	EoTypeStaticArray     EoType = C.EOLIAN_TYPE_STATIC_ARRAY
	EoTypeTerminatedArray EoType = C.EOLIAN_TYPE_TERMINATED_ARRAY
	EoTypeUndefined       EoType = C.EOLIAN_TYPE_UNDEFINED
)

type ExpressionType int

const (
	ExpressionTypeUnknown ExpressionType = C.EOLIAN_EXPR_UNKNOWN
	ExpressionTypeInt     ExpressionType = C.EOLIAN_EXPR_INT
	ExpressionTypeUint    ExpressionType = C.EOLIAN_EXPR_UINT
	ExpressionTypeLong    ExpressionType = C.EOLIAN_EXPR_LONG
	ExpressionTypeUlong   ExpressionType = C.EOLIAN_EXPR_ULONG
	ExpressionTypeLlong   ExpressionType = C.EOLIAN_EXPR_LLONG
	ExpressionTypeUllong  ExpressionType = C.EOLIAN_EXPR_ULLONG
	ExpressionTypeFloat   ExpressionType = C.EOLIAN_EXPR_FLOAT
	ExpressionTypeDouble  ExpressionType = C.EOLIAN_EXPR_DOUBLE
	ExpressionTypeString  ExpressionType = C.EOLIAN_EXPR_STRING
	ExpressionTypeChar    ExpressionType = C.EOLIAN_EXPR_CHAR
	ExpressionTypeNull    ExpressionType = C.EOLIAN_EXPR_NULL
	ExpressionTypeBool    ExpressionType = C.EOLIAN_EXPR_BOOL
	ExpressionTypeName    ExpressionType = C.EOLIAN_EXPR_NAME
	ExpressionTypeUnary   ExpressionType = C.EOLIAN_EXPR_UNARY
	ExpressionTypeBinary  ExpressionType = C.EOLIAN_EXPR_BINARY
)

type ExpressionMask int

const (
	ExpressionMaskSint   ExpressionMask = C.EOLIAN_MASK_SINT
	ExpressionMaskUint   ExpressionMask = C.EOLIAN_MASK_UINT
	ExpressionMaskInt    ExpressionMask = C.EOLIAN_MASK_INT
	ExpressionMaskFloat  ExpressionMask = C.EOLIAN_MASK_FLOAT
	ExpressionMaskBool   ExpressionMask = C.EOLIAN_MASK_BOOL
	ExpressionMaskString ExpressionMask = C.EOLIAN_MASK_STRING
	ExpressionMaskChar   ExpressionMask = C.EOLIAN_MASK_CHAR
	ExpressionMaskNull   ExpressionMask = C.EOLIAN_MASK_NULL
	// not in 1.18.4
	//ExpressionMaskSigned ExpressionMask = C.EOLIAN_MASK_SIGNED
	ExpressionMaskNumber ExpressionMask = C.EOLIAN_MASK_NUMBER
	ExpressionMaskAll    ExpressionMask = C.EOLIAN_MASK_ALL
)

type VariableType int

const (
	VariableTypeUnknown  VariableType = C.EOLIAN_VAR_UNKNOWN
	VariableTypeConstant VariableType = C.EOLIAN_VAR_CONSTANT
	VariableTypeGlobal   VariableType = C.EOLIAN_VAR_GLOBAL
)

type BinaryOperator int

const (
	BinaryOperatorInvalid BinaryOperator = C.EOLIAN_BINOP_INVALID
	BinaryOperatorAdd     BinaryOperator = C.EOLIAN_BINOP_ADD  /* + int, float */
	BinaryOperatorSub     BinaryOperator = C.EOLIAN_BINOP_SUB  /* - int, float */
	BinaryOperatorMul     BinaryOperator = C.EOLIAN_BINOP_MUL  /* * int, float */
	BinaryOperatorDiv     BinaryOperator = C.EOLIAN_BINOP_DIV  /* / int, float */
	BinaryOperatorMod     BinaryOperator = C.EOLIAN_BINOP_MOD  /* % int */
	BinaryOperatorEq      BinaryOperator = C.EOLIAN_BINOP_EQ   /* == all types */
	BinaryOperatorNq      BinaryOperator = C.EOLIAN_BINOP_NQ   /* != all types */
	BinaryOperatorGt      BinaryOperator = C.EOLIAN_BINOP_GT   /* >  int, float */
	BinaryOperatorLt      BinaryOperator = C.EOLIAN_BINOP_LT   /* <  int, float */
	BinaryOperatorGe      BinaryOperator = C.EOLIAN_BINOP_GE   /* >= int, float */
	BinaryOperatorLe      BinaryOperator = C.EOLIAN_BINOP_LE   /* <= int, float */
	BinaryOperatorAnd     BinaryOperator = C.EOLIAN_BINOP_AND  /* && all types */
	BinaryOperatorOr      BinaryOperator = C.EOLIAN_BINOP_OR   /* || all types */
	BinaryOperatorBand    BinaryOperator = C.EOLIAN_BINOP_BAND /* &  int */
	BinaryOperatorBor     BinaryOperator = C.EOLIAN_BINOP_BOR  /* |  int */
	BinaryOperatorBxor    BinaryOperator = C.EOLIAN_BINOP_BXOR /* ^  int */
	BinaryOperatorLsh     BinaryOperator = C.EOLIAN_BINOP_LSH  /* << int */
	BinaryOperatorRsh     BinaryOperator = C.EOLIAN_BINOP_RSH  /* >> int */
)

type UnaryOperator int

const (
	UnaryOperatorInvalid UnaryOperator = C.EOLIAN_UNOP_INVALID
	UnaryOperatorUnm     UnaryOperator = C.EOLIAN_UNOP_UNM  /* - sint */
	UnaryOperatorUnp     UnaryOperator = C.EOLIAN_UNOP_UNP  /* + sint */
	UnaryOperatorNot     UnaryOperator = C.EOLIAN_UNOP_NOT  /* ! int, float, bool */
	UnaryOperatorBnot    UnaryOperator = C.EOLIAN_UNOP_BNOT /* ~ int */
)

type DeclarationType int

const (
	DeclarationTypeUnknown DeclarationType = C.EOLIAN_DECL_UNKNOWN
	DeclarationTypeClass   DeclarationType = C.EOLIAN_DECL_CLASS
	DeclarationTypeAlias   DeclarationType = C.EOLIAN_DECL_ALIAS
	DeclarationTypeStruct  DeclarationType = C.EOLIAN_DECL_STRUCT
	DeclarationTypeEnum    DeclarationType = C.EOLIAN_DECL_ENUM
	DeclarationTypeVar     DeclarationType = C.EOLIAN_DECL_VAR
)

/* not in 1.18.4
type DocTokenType int

const (
	DocTokenTypeUnknown         DocTokenType = C.EOLIAN_DOC_TOKEN_UNKNOWN
	DocTokenTypeText            DocTokenType = C.EOLIAN_DOC_TOKEN_TEXT
	DocTokenTypeRef             DocTokenType = C.EOLIAN_DOC_TOKEN_REF
	DocTokenTypeNote            DocTokenType = C.EOLIAN_DOC_TOKEN_MARK_NOTE
	DocTokenTypeMarkWarning     DocTokenType = C.EOLIAN_DOC_TOKEN_MARK_WARNING
	DocTokenTypeMarkRemark      DocTokenType = C.EOLIAN_DOC_TOKEN_MARK_REMARK
	DocTokenTypeMarkTodo        DocTokenType = C.EOLIAN_DOC_TOKEN_MARK_TODO
	DocTokenTypeMarkupMonospace DocTokenType = C.EOLIAN_DOC_TOKEN_MARKUP_MONOSPACE
)
*/

/* not in 1.18.4
type DocRefType int

const (
	DocRefTypeInvalid     DocRefType = C.EOLIAN_DOC_REF_INVALID
	DocRefTypeClass       DocRefType = C.EOLIAN_DOC_REF_CLASS
	DocRefTypeFunc        DocRefType = C.EOLIAN_DOC_REF_FUNC
	DocRefTypeEvent       DocRefType = C.EOLIAN_DOC_REF_EVENT
	DocRefTypeAlias       DocRefType = C.EOLIAN_DOC_REF_ALIAS
	DocRefTypeStruct      DocRefType = C.EOLIAN_DOC_REF_STRUCT
	DocRefTypeStructField DocRefType = C.EOLIAN_DOC_REF_STRUCT_FIELD
	DocRefTypeEnum        DocRefType = C.EOLIAN_DOC_REF_ENUM
	DocRefTypeEnumField   DocRefType = C.EOLIAN_DOC_REF_ENUM_FIELD
	DocRefTypeVar         DocRefType = C.EOLIAN_DOC_REF_VAR
)

*/
