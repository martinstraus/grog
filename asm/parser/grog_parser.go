// Code generated from Grog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grog
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 197,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 6, 2, 52, 10, 2, 13, 2, 14, 2, 53, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 5, 3, 70, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 78,
	10, 4, 3, 5, 3, 5, 5, 5, 82, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 88, 10,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 100,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 152, 10, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 5, 21, 163,
	10, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 179, 10, 22, 3, 22, 3, 22, 3,
	22, 5, 22, 184, 10, 22, 3, 23, 3, 23, 5, 23, 188, 10, 23, 3, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 2, 26, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	2, 2, 2, 217, 2, 51, 3, 2, 2, 2, 4, 69, 3, 2, 2, 2, 6, 71, 3, 2, 2, 2,
	8, 81, 3, 2, 2, 2, 10, 99, 3, 2, 2, 2, 12, 101, 3, 2, 2, 2, 14, 105, 3,
	2, 2, 2, 16, 109, 3, 2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 117, 3, 2, 2, 2,
	22, 121, 3, 2, 2, 2, 24, 125, 3, 2, 2, 2, 26, 129, 3, 2, 2, 2, 28, 133,
	3, 2, 2, 2, 30, 137, 3, 2, 2, 2, 32, 141, 3, 2, 2, 2, 34, 144, 3, 2, 2,
	2, 36, 151, 3, 2, 2, 2, 38, 156, 3, 2, 2, 2, 40, 162, 3, 2, 2, 2, 42, 178,
	3, 2, 2, 2, 44, 187, 3, 2, 2, 2, 46, 192, 3, 2, 2, 2, 48, 194, 3, 2, 2,
	2, 50, 52, 5, 4, 3, 2, 51, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 7, 2, 2, 3,
	56, 3, 3, 2, 2, 2, 57, 70, 5, 32, 17, 2, 58, 70, 5, 34, 18, 2, 59, 70,
	5, 36, 19, 2, 60, 70, 5, 38, 20, 2, 61, 70, 5, 40, 21, 2, 62, 70, 5, 6,
	4, 2, 63, 70, 5, 8, 5, 2, 64, 70, 5, 10, 6, 2, 65, 70, 5, 42, 22, 2, 66,
	70, 5, 44, 23, 2, 67, 70, 5, 46, 24, 2, 68, 70, 5, 48, 25, 2, 69, 57, 3,
	2, 2, 2, 69, 58, 3, 2, 2, 2, 69, 59, 3, 2, 2, 2, 69, 60, 3, 2, 2, 2, 69,
	61, 3, 2, 2, 2, 69, 62, 3, 2, 2, 2, 69, 63, 3, 2, 2, 2, 69, 64, 3, 2, 2,
	2, 69, 65, 3, 2, 2, 2, 69, 66, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 68,
	3, 2, 2, 2, 70, 5, 3, 2, 2, 2, 71, 72, 7, 8, 2, 2, 72, 77, 7, 39, 2, 2,
	73, 78, 7, 37, 2, 2, 74, 78, 7, 41, 2, 2, 75, 78, 7, 42, 2, 2, 76, 78,
	7, 43, 2, 2, 77, 73, 3, 2, 2, 2, 77, 74, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2,
	77, 76, 3, 2, 2, 2, 78, 7, 3, 2, 2, 2, 79, 82, 7, 39, 2, 2, 80, 82, 7,
	37, 2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83,
	87, 7, 3, 2, 2, 84, 88, 7, 41, 2, 2, 85, 88, 7, 42, 2, 2, 86, 88, 7, 43,
	2, 2, 87, 84, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 9,
	3, 2, 2, 2, 89, 100, 5, 12, 7, 2, 90, 100, 5, 14, 8, 2, 91, 100, 5, 16,
	9, 2, 92, 100, 5, 18, 10, 2, 93, 100, 5, 20, 11, 2, 94, 100, 5, 22, 12,
	2, 95, 100, 5, 24, 13, 2, 96, 100, 5, 26, 14, 2, 97, 100, 5, 28, 15, 2,
	98, 100, 5, 30, 16, 2, 99, 89, 3, 2, 2, 2, 99, 90, 3, 2, 2, 2, 99, 91,
	3, 2, 2, 2, 99, 92, 3, 2, 2, 2, 99, 93, 3, 2, 2, 2, 99, 94, 3, 2, 2, 2,
	99, 95, 3, 2, 2, 2, 99, 96, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 98, 3,
	2, 2, 2, 100, 11, 3, 2, 2, 2, 101, 102, 7, 10, 2, 2, 102, 103, 7, 39, 2,
	2, 103, 104, 7, 39, 2, 2, 104, 13, 3, 2, 2, 2, 105, 106, 7, 10, 2, 2, 106,
	107, 7, 41, 2, 2, 107, 108, 7, 41, 2, 2, 108, 15, 3, 2, 2, 2, 109, 110,
	7, 10, 2, 2, 110, 111, 7, 41, 2, 2, 111, 112, 7, 42, 2, 2, 112, 17, 3,
	2, 2, 2, 113, 114, 7, 10, 2, 2, 114, 115, 7, 41, 2, 2, 115, 116, 7, 43,
	2, 2, 116, 19, 3, 2, 2, 2, 117, 118, 7, 10, 2, 2, 118, 119, 7, 42, 2, 2,
	119, 120, 7, 41, 2, 2, 120, 21, 3, 2, 2, 2, 121, 122, 7, 10, 2, 2, 122,
	123, 7, 42, 2, 2, 123, 124, 7, 42, 2, 2, 124, 23, 3, 2, 2, 2, 125, 126,
	7, 10, 2, 2, 126, 127, 7, 42, 2, 2, 127, 128, 7, 43, 2, 2, 128, 25, 3,
	2, 2, 2, 129, 130, 7, 10, 2, 2, 130, 131, 7, 43, 2, 2, 131, 132, 7, 41,
	2, 2, 132, 27, 3, 2, 2, 2, 133, 134, 7, 10, 2, 2, 134, 135, 7, 43, 2, 2,
	135, 136, 7, 42, 2, 2, 136, 29, 3, 2, 2, 2, 137, 138, 7, 10, 2, 2, 138,
	139, 7, 43, 2, 2, 139, 140, 7, 43, 2, 2, 140, 31, 3, 2, 2, 2, 141, 142,
	7, 11, 2, 2, 142, 143, 7, 39, 2, 2, 143, 33, 3, 2, 2, 2, 144, 145, 7, 12,
	2, 2, 145, 146, 7, 39, 2, 2, 146, 35, 3, 2, 2, 2, 147, 152, 7, 13, 2, 2,
	148, 152, 7, 14, 2, 2, 149, 152, 7, 16, 2, 2, 150, 152, 7, 15, 2, 2, 151,
	147, 3, 2, 2, 2, 151, 148, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 150,
	3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 7, 39, 2, 2, 154, 155, 7, 39,
	2, 2, 155, 37, 3, 2, 2, 2, 156, 157, 7, 28, 2, 2, 157, 158, 7, 39, 2, 2,
	158, 39, 3, 2, 2, 2, 159, 163, 7, 29, 2, 2, 160, 163, 7, 30, 2, 2, 161,
	163, 7, 31, 2, 2, 162, 159, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 161,
	3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 7, 39, 2, 2, 165, 166, 7, 39,
	2, 2, 166, 41, 3, 2, 2, 2, 167, 179, 7, 17, 2, 2, 168, 179, 7, 18, 2, 2,
	169, 179, 7, 19, 2, 2, 170, 179, 7, 20, 2, 2, 171, 179, 7, 21, 2, 2, 172,
	179, 7, 22, 2, 2, 173, 179, 7, 23, 2, 2, 174, 179, 7, 24, 2, 2, 175, 179,
	7, 25, 2, 2, 176, 179, 7, 26, 2, 2, 177, 179, 7, 27, 2, 2, 178, 167, 3,
	2, 2, 2, 178, 168, 3, 2, 2, 2, 178, 169, 3, 2, 2, 2, 178, 170, 3, 2, 2,
	2, 178, 171, 3, 2, 2, 2, 178, 172, 3, 2, 2, 2, 178, 173, 3, 2, 2, 2, 178,
	174, 3, 2, 2, 2, 178, 175, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 177,
	3, 2, 2, 2, 179, 183, 3, 2, 2, 2, 180, 184, 7, 41, 2, 2, 181, 184, 7, 42,
	2, 2, 182, 184, 7, 43, 2, 2, 183, 180, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2,
	183, 182, 3, 2, 2, 2, 184, 43, 3, 2, 2, 2, 185, 188, 7, 32, 2, 2, 186,
	188, 7, 33, 2, 2, 187, 185, 3, 2, 2, 2, 187, 186, 3, 2, 2, 2, 188, 189,
	3, 2, 2, 2, 189, 190, 7, 39, 2, 2, 190, 191, 7, 40, 2, 2, 191, 45, 3, 2,
	2, 2, 192, 193, 7, 34, 2, 2, 193, 47, 3, 2, 2, 2, 194, 195, 7, 35, 2, 2,
	195, 49, 3, 2, 2, 2, 13, 53, 69, 77, 81, 87, 99, 151, 162, 178, 183, 187,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'->'", "", "", "", "", "'load'", "'move'", "'copy'", "'increment'",
	"'decrement'", "'add'", "'subtract'", "'divide'", "'multiply'", "'jump'",
	"'je'", "'jne'", "'jg'", "'jng'", "'jge'", "'jnge'", "'jl'", "'jnl'", "'jle'",
	"'jnle'", "'not'", "'and'", "'or'", "'xor'", "'input'", "'output'", "'stop'",
	"'wait'",
}
var symbolicNames = []string{
	"", "", "WHITESPACE", "WS", "COMMENT", "LINE_COMMENT", "LOAD", "MOVE",
	"COPY", "INCREMENT", "DECREMENT", "ADD", "SUBTRACT", "DIVIDE", "MULTIPLY",
	"JUMP", "JUMP_IF_EQUAL", "JUMP_IF_NOT_EQUAL", "JUMP_IF_GREATER", "JUMP_IF_NOT_GREATER",
	"JUMP_IF_GREATER_OR_EQUAL", "JUMP_IF_NOT_GREATER_OR_EQUAL", "JUMP_IF_LESS",
	"JUMP_IF_NOT_LESS", "JUMP_IF_LESS_OR_EQUAL", "JUMP_IF_NOT_LESS_OR_EQUAL",
	"NOT", "AND", "OR", "XOR", "INPUT", "OUTPUT", "STOP", "WAIT", "HEX_DIGIT",
	"HEXA_BYTE", "WORD", "REGISTER", "DEVICE", "ABSOLUTE_ADDRESS", "OFFSET_ADDRESS",
	"POINTER_ADDRESS",
}

var ruleNames = []string{
	"program", "instruction", "load", "store", "copyValue", "copyRegister",
	"copyAbsoluteToAbsolute", "copyOffsetToAbsolute", "copyPointerToAbsolute",
	"copyAbsoluteToOffset", "copyOffsetToOffset", "copyPointerToOffset", "copyAbsoluteToPointer",
	"copyOffsetToPointer", "copyPointerToPointer", "increment", "decrement",
	"arithmeticOperation", "unaryBooleanOperation", "binaryBooleanOperation",
	"jump", "io", "stop", "wait",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrogParser struct {
	*antlr.BaseParser
}

func NewGrogParser(input antlr.TokenStream) *GrogParser {
	this := new(GrogParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Grog.g4"

	return this
}

// GrogParser tokens.
const (
	GrogParserEOF                          = antlr.TokenEOF
	GrogParserT__0                         = 1
	GrogParserWHITESPACE                   = 2
	GrogParserWS                           = 3
	GrogParserCOMMENT                      = 4
	GrogParserLINE_COMMENT                 = 5
	GrogParserLOAD                         = 6
	GrogParserMOVE                         = 7
	GrogParserCOPY                         = 8
	GrogParserINCREMENT                    = 9
	GrogParserDECREMENT                    = 10
	GrogParserADD                          = 11
	GrogParserSUBTRACT                     = 12
	GrogParserDIVIDE                       = 13
	GrogParserMULTIPLY                     = 14
	GrogParserJUMP                         = 15
	GrogParserJUMP_IF_EQUAL                = 16
	GrogParserJUMP_IF_NOT_EQUAL            = 17
	GrogParserJUMP_IF_GREATER              = 18
	GrogParserJUMP_IF_NOT_GREATER          = 19
	GrogParserJUMP_IF_GREATER_OR_EQUAL     = 20
	GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL = 21
	GrogParserJUMP_IF_LESS                 = 22
	GrogParserJUMP_IF_NOT_LESS             = 23
	GrogParserJUMP_IF_LESS_OR_EQUAL        = 24
	GrogParserJUMP_IF_NOT_LESS_OR_EQUAL    = 25
	GrogParserNOT                          = 26
	GrogParserAND                          = 27
	GrogParserOR                           = 28
	GrogParserXOR                          = 29
	GrogParserINPUT                        = 30
	GrogParserOUTPUT                       = 31
	GrogParserSTOP                         = 32
	GrogParserWAIT                         = 33
	GrogParserHEX_DIGIT                    = 34
	GrogParserHEXA_BYTE                    = 35
	GrogParserWORD                         = 36
	GrogParserREGISTER                     = 37
	GrogParserDEVICE                       = 38
	GrogParserABSOLUTE_ADDRESS             = 39
	GrogParserOFFSET_ADDRESS               = 40
	GrogParserPOINTER_ADDRESS              = 41
)

// GrogParser rules.
const (
	GrogParserRULE_program                = 0
	GrogParserRULE_instruction            = 1
	GrogParserRULE_load                   = 2
	GrogParserRULE_store                  = 3
	GrogParserRULE_copyValue              = 4
	GrogParserRULE_copyRegister           = 5
	GrogParserRULE_copyAbsoluteToAbsolute = 6
	GrogParserRULE_copyOffsetToAbsolute   = 7
	GrogParserRULE_copyPointerToAbsolute  = 8
	GrogParserRULE_copyAbsoluteToOffset   = 9
	GrogParserRULE_copyOffsetToOffset     = 10
	GrogParserRULE_copyPointerToOffset    = 11
	GrogParserRULE_copyAbsoluteToPointer  = 12
	GrogParserRULE_copyOffsetToPointer    = 13
	GrogParserRULE_copyPointerToPointer   = 14
	GrogParserRULE_increment              = 15
	GrogParserRULE_decrement              = 16
	GrogParserRULE_arithmeticOperation    = 17
	GrogParserRULE_unaryBooleanOperation  = 18
	GrogParserRULE_binaryBooleanOperation = 19
	GrogParserRULE_jump                   = 20
	GrogParserRULE_io                     = 21
	GrogParserRULE_stop                   = 22
	GrogParserRULE_wait                   = 23
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrogParserEOF, 0)
}

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *GrogParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrogParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(GrogParserLOAD-6))|(1<<(GrogParserCOPY-6))|(1<<(GrogParserINCREMENT-6))|(1<<(GrogParserDECREMENT-6))|(1<<(GrogParserADD-6))|(1<<(GrogParserSUBTRACT-6))|(1<<(GrogParserDIVIDE-6))|(1<<(GrogParserMULTIPLY-6))|(1<<(GrogParserJUMP-6))|(1<<(GrogParserJUMP_IF_EQUAL-6))|(1<<(GrogParserJUMP_IF_NOT_EQUAL-6))|(1<<(GrogParserJUMP_IF_GREATER-6))|(1<<(GrogParserJUMP_IF_NOT_GREATER-6))|(1<<(GrogParserJUMP_IF_GREATER_OR_EQUAL-6))|(1<<(GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL-6))|(1<<(GrogParserJUMP_IF_LESS-6))|(1<<(GrogParserJUMP_IF_NOT_LESS-6))|(1<<(GrogParserJUMP_IF_LESS_OR_EQUAL-6))|(1<<(GrogParserJUMP_IF_NOT_LESS_OR_EQUAL-6))|(1<<(GrogParserNOT-6))|(1<<(GrogParserAND-6))|(1<<(GrogParserOR-6))|(1<<(GrogParserXOR-6))|(1<<(GrogParserINPUT-6))|(1<<(GrogParserOUTPUT-6))|(1<<(GrogParserSTOP-6))|(1<<(GrogParserWAIT-6))|(1<<(GrogParserHEXA_BYTE-6))|(1<<(GrogParserREGISTER-6)))) != 0) {
		{
			p.SetState(48)
			p.Instruction()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(GrogParserEOF)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Increment() IIncrementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncrementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncrementContext)
}

func (s *InstructionContext) Decrement() IDecrementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecrementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecrementContext)
}

func (s *InstructionContext) ArithmeticOperation() IArithmeticOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArithmeticOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArithmeticOperationContext)
}

func (s *InstructionContext) UnaryBooleanOperation() IUnaryBooleanOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryBooleanOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryBooleanOperationContext)
}

func (s *InstructionContext) BinaryBooleanOperation() IBinaryBooleanOperationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryBooleanOperationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryBooleanOperationContext)
}

func (s *InstructionContext) Load() ILoadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoadContext)
}

func (s *InstructionContext) Store() IStoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStoreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStoreContext)
}

func (s *InstructionContext) CopyValue() ICopyValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyValueContext)
}

func (s *InstructionContext) Jump() IJumpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJumpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJumpContext)
}

func (s *InstructionContext) Io() IIoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIoContext)
}

func (s *InstructionContext) Stop() IStopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStopContext)
}

func (s *InstructionContext) Wait() IWaitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWaitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWaitContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *GrogParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrogParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserINCREMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Increment()
		}

	case GrogParserDECREMENT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Decrement()
		}

	case GrogParserADD, GrogParserSUBTRACT, GrogParserDIVIDE, GrogParserMULTIPLY:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.ArithmeticOperation()
		}

	case GrogParserNOT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.UnaryBooleanOperation()
		}

	case GrogParserAND, GrogParserOR, GrogParserXOR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.BinaryBooleanOperation()
		}

	case GrogParserLOAD:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Load()
		}

	case GrogParserHEXA_BYTE, GrogParserREGISTER:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Store()
		}

	case GrogParserCOPY:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.CopyValue()
		}

	case GrogParserJUMP, GrogParserJUMP_IF_EQUAL, GrogParserJUMP_IF_NOT_EQUAL, GrogParserJUMP_IF_GREATER, GrogParserJUMP_IF_NOT_GREATER, GrogParserJUMP_IF_GREATER_OR_EQUAL, GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL, GrogParserJUMP_IF_LESS, GrogParserJUMP_IF_NOT_LESS, GrogParserJUMP_IF_LESS_OR_EQUAL, GrogParserJUMP_IF_NOT_LESS_OR_EQUAL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.Jump()
		}

	case GrogParserINPUT, GrogParserOUTPUT:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(64)
			p.Io()
		}

	case GrogParserSTOP:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(65)
			p.Stop()
		}

	case GrogParserWAIT:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(66)
			p.Wait()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoadContext is an interface to support dynamic dispatch.
type ILoadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// GetValue returns the Value token.
	GetValue() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// SetValue sets the Value token.
	SetValue(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsLoadContext differentiates from other interfaces.
	IsLoadContext()
}

type LoadContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
	Value    antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyLoadContext() *LoadContext {
	var p = new(LoadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_load
	return p
}

func (*LoadContext) IsLoadContext() {}

func NewLoadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadContext {
	var p = new(LoadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_load

	return p
}

func (s *LoadContext) GetParser() antlr.Parser { return s.parser }

func (s *LoadContext) GetRegister() antlr.Token { return s.Register }

func (s *LoadContext) GetValue() antlr.Token { return s.Value }

func (s *LoadContext) GetAddress() antlr.Token { return s.Address }

func (s *LoadContext) GetOffset() antlr.Token { return s.Offset }

func (s *LoadContext) GetPointer() antlr.Token { return s.Pointer }

func (s *LoadContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *LoadContext) SetValue(v antlr.Token) { s.Value = v }

func (s *LoadContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *LoadContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *LoadContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *LoadContext) LOAD() antlr.TerminalNode {
	return s.GetToken(GrogParserLOAD, 0)
}

func (s *LoadContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *LoadContext) HEXA_BYTE() antlr.TerminalNode {
	return s.GetToken(GrogParserHEXA_BYTE, 0)
}

func (s *LoadContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *LoadContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *LoadContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *LoadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterLoad(s)
	}
}

func (s *LoadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitLoad(s)
	}
}

func (p *GrogParser) Load() (localctx ILoadContext) {
	localctx = NewLoadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrogParserRULE_load)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(GrogParserLOAD)
	}
	{
		p.SetState(70)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*LoadContext).Register = _m
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserHEXA_BYTE:
		{
			p.SetState(71)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*LoadContext).Value = _m
		}

	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(72)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*LoadContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(73)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*LoadContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(74)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*LoadContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStoreContext is an interface to support dynamic dispatch.
type IStoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// GetValue returns the Value token.
	GetValue() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// SetValue sets the Value token.
	SetValue(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsStoreContext differentiates from other interfaces.
	IsStoreContext()
}

type StoreContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
	Value    antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyStoreContext() *StoreContext {
	var p = new(StoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_store
	return p
}

func (*StoreContext) IsStoreContext() {}

func NewStoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreContext {
	var p = new(StoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_store

	return p
}

func (s *StoreContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreContext) GetRegister() antlr.Token { return s.Register }

func (s *StoreContext) GetValue() antlr.Token { return s.Value }

func (s *StoreContext) GetAddress() antlr.Token { return s.Address }

func (s *StoreContext) GetOffset() antlr.Token { return s.Offset }

func (s *StoreContext) GetPointer() antlr.Token { return s.Pointer }

func (s *StoreContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *StoreContext) SetValue(v antlr.Token) { s.Value = v }

func (s *StoreContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *StoreContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *StoreContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *StoreContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *StoreContext) HEXA_BYTE() antlr.TerminalNode {
	return s.GetToken(GrogParserHEXA_BYTE, 0)
}

func (s *StoreContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *StoreContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *StoreContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *StoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterStore(s)
	}
}

func (s *StoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitStore(s)
	}
}

func (p *GrogParser) Store() (localctx IStoreContext) {
	localctx = NewStoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrogParserRULE_store)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserREGISTER:
		{
			p.SetState(77)

			var _m = p.Match(GrogParserREGISTER)

			localctx.(*StoreContext).Register = _m
		}

	case GrogParserHEXA_BYTE:
		{
			p.SetState(78)

			var _m = p.Match(GrogParserHEXA_BYTE)

			localctx.(*StoreContext).Value = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(81)
		p.Match(GrogParserT__0)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(82)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*StoreContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(83)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*StoreContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(84)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*StoreContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICopyValueContext is an interface to support dynamic dispatch.
type ICopyValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopyValueContext differentiates from other interfaces.
	IsCopyValueContext()
}

type CopyValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopyValueContext() *CopyValueContext {
	var p = new(CopyValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyValue
	return p
}

func (*CopyValueContext) IsCopyValueContext() {}

func NewCopyValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyValueContext {
	var p = new(CopyValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyValue

	return p
}

func (s *CopyValueContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyValueContext) CopyRegister() ICopyRegisterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyRegisterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyRegisterContext)
}

func (s *CopyValueContext) CopyAbsoluteToAbsolute() ICopyAbsoluteToAbsoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyAbsoluteToAbsoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyAbsoluteToAbsoluteContext)
}

func (s *CopyValueContext) CopyOffsetToAbsolute() ICopyOffsetToAbsoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyOffsetToAbsoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyOffsetToAbsoluteContext)
}

func (s *CopyValueContext) CopyPointerToAbsolute() ICopyPointerToAbsoluteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyPointerToAbsoluteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyPointerToAbsoluteContext)
}

func (s *CopyValueContext) CopyAbsoluteToOffset() ICopyAbsoluteToOffsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyAbsoluteToOffsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyAbsoluteToOffsetContext)
}

func (s *CopyValueContext) CopyOffsetToOffset() ICopyOffsetToOffsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyOffsetToOffsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyOffsetToOffsetContext)
}

func (s *CopyValueContext) CopyPointerToOffset() ICopyPointerToOffsetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyPointerToOffsetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyPointerToOffsetContext)
}

func (s *CopyValueContext) CopyAbsoluteToPointer() ICopyAbsoluteToPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyAbsoluteToPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyAbsoluteToPointerContext)
}

func (s *CopyValueContext) CopyOffsetToPointer() ICopyOffsetToPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyOffsetToPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyOffsetToPointerContext)
}

func (s *CopyValueContext) CopyPointerToPointer() ICopyPointerToPointerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyPointerToPointerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICopyPointerToPointerContext)
}

func (s *CopyValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyValue(s)
	}
}

func (s *CopyValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyValue(s)
	}
}

func (p *GrogParser) CopyValue() (localctx ICopyValueContext) {
	localctx = NewCopyValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrogParserRULE_copyValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(87)
			p.CopyRegister()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)
			p.CopyAbsoluteToAbsolute()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.CopyOffsetToAbsolute()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(90)
			p.CopyPointerToAbsolute()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(91)
			p.CopyAbsoluteToOffset()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(92)
			p.CopyOffsetToOffset()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(93)
			p.CopyPointerToOffset()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(94)
			p.CopyAbsoluteToPointer()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(95)
			p.CopyOffsetToPointer()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(96)
			p.CopyPointerToPointer()
		}

	}

	return localctx
}

// ICopyRegisterContext is an interface to support dynamic dispatch.
type ICopyRegisterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestinationRegister returns the DestinationRegister token.
	GetDestinationRegister() antlr.Token

	// GetSourceRegister returns the SourceRegister token.
	GetSourceRegister() antlr.Token

	// SetDestinationRegister sets the DestinationRegister token.
	SetDestinationRegister(antlr.Token)

	// SetSourceRegister sets the SourceRegister token.
	SetSourceRegister(antlr.Token)

	// IsCopyRegisterContext differentiates from other interfaces.
	IsCopyRegisterContext()
}

type CopyRegisterContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	DestinationRegister antlr.Token
	SourceRegister      antlr.Token
}

func NewEmptyCopyRegisterContext() *CopyRegisterContext {
	var p = new(CopyRegisterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyRegister
	return p
}

func (*CopyRegisterContext) IsCopyRegisterContext() {}

func NewCopyRegisterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyRegisterContext {
	var p = new(CopyRegisterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyRegister

	return p
}

func (s *CopyRegisterContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyRegisterContext) GetDestinationRegister() antlr.Token { return s.DestinationRegister }

func (s *CopyRegisterContext) GetSourceRegister() antlr.Token { return s.SourceRegister }

func (s *CopyRegisterContext) SetDestinationRegister(v antlr.Token) { s.DestinationRegister = v }

func (s *CopyRegisterContext) SetSourceRegister(v antlr.Token) { s.SourceRegister = v }

func (s *CopyRegisterContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyRegisterContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *CopyRegisterContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *CopyRegisterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyRegisterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyRegisterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyRegister(s)
	}
}

func (s *CopyRegisterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyRegister(s)
	}
}

func (p *GrogParser) CopyRegister() (localctx ICopyRegisterContext) {
	localctx = NewCopyRegisterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrogParserRULE_copyRegister)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(100)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*CopyRegisterContext).DestinationRegister = _m
	}
	{
		p.SetState(101)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*CopyRegisterContext).SourceRegister = _m
	}

	return localctx
}

// ICopyAbsoluteToAbsoluteContext is an interface to support dynamic dispatch.
type ICopyAbsoluteToAbsoluteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyAbsoluteToAbsoluteContext differentiates from other interfaces.
	IsCopyAbsoluteToAbsoluteContext()
}

type CopyAbsoluteToAbsoluteContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyAbsoluteToAbsoluteContext() *CopyAbsoluteToAbsoluteContext {
	var p = new(CopyAbsoluteToAbsoluteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyAbsoluteToAbsolute
	return p
}

func (*CopyAbsoluteToAbsoluteContext) IsCopyAbsoluteToAbsoluteContext() {}

func NewCopyAbsoluteToAbsoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyAbsoluteToAbsoluteContext {
	var p = new(CopyAbsoluteToAbsoluteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyAbsoluteToAbsolute

	return p
}

func (s *CopyAbsoluteToAbsoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyAbsoluteToAbsoluteContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyAbsoluteToAbsoluteContext) GetSource() antlr.Token { return s.Source }

func (s *CopyAbsoluteToAbsoluteContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyAbsoluteToAbsoluteContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyAbsoluteToAbsoluteContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyAbsoluteToAbsoluteContext) AllABSOLUTE_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserABSOLUTE_ADDRESS)
}

func (s *CopyAbsoluteToAbsoluteContext) ABSOLUTE_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, i)
}

func (s *CopyAbsoluteToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToAbsoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyAbsoluteToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToAbsolute(s)
	}
}

func (s *CopyAbsoluteToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToAbsolute(s)
	}
}

func (p *GrogParser) CopyAbsoluteToAbsolute() (localctx ICopyAbsoluteToAbsoluteContext) {
	localctx = NewCopyAbsoluteToAbsoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrogParserRULE_copyAbsoluteToAbsolute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(104)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyAbsoluteToAbsoluteContext).Destination = _m
	}
	{
		p.SetState(105)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyAbsoluteToAbsoluteContext).Source = _m
	}

	return localctx
}

// ICopyOffsetToAbsoluteContext is an interface to support dynamic dispatch.
type ICopyOffsetToAbsoluteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyOffsetToAbsoluteContext differentiates from other interfaces.
	IsCopyOffsetToAbsoluteContext()
}

type CopyOffsetToAbsoluteContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyOffsetToAbsoluteContext() *CopyOffsetToAbsoluteContext {
	var p = new(CopyOffsetToAbsoluteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyOffsetToAbsolute
	return p
}

func (*CopyOffsetToAbsoluteContext) IsCopyOffsetToAbsoluteContext() {}

func NewCopyOffsetToAbsoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyOffsetToAbsoluteContext {
	var p = new(CopyOffsetToAbsoluteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyOffsetToAbsolute

	return p
}

func (s *CopyOffsetToAbsoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyOffsetToAbsoluteContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyOffsetToAbsoluteContext) GetSource() antlr.Token { return s.Source }

func (s *CopyOffsetToAbsoluteContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyOffsetToAbsoluteContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyOffsetToAbsoluteContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyOffsetToAbsoluteContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyOffsetToAbsoluteContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyOffsetToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToAbsoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyOffsetToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToAbsolute(s)
	}
}

func (s *CopyOffsetToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToAbsolute(s)
	}
}

func (p *GrogParser) CopyOffsetToAbsolute() (localctx ICopyOffsetToAbsoluteContext) {
	localctx = NewCopyOffsetToAbsoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrogParserRULE_copyOffsetToAbsolute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(108)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyOffsetToAbsoluteContext).Destination = _m
	}
	{
		p.SetState(109)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyOffsetToAbsoluteContext).Source = _m
	}

	return localctx
}

// ICopyPointerToAbsoluteContext is an interface to support dynamic dispatch.
type ICopyPointerToAbsoluteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyPointerToAbsoluteContext differentiates from other interfaces.
	IsCopyPointerToAbsoluteContext()
}

type CopyPointerToAbsoluteContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyPointerToAbsoluteContext() *CopyPointerToAbsoluteContext {
	var p = new(CopyPointerToAbsoluteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyPointerToAbsolute
	return p
}

func (*CopyPointerToAbsoluteContext) IsCopyPointerToAbsoluteContext() {}

func NewCopyPointerToAbsoluteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyPointerToAbsoluteContext {
	var p = new(CopyPointerToAbsoluteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyPointerToAbsolute

	return p
}

func (s *CopyPointerToAbsoluteContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyPointerToAbsoluteContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyPointerToAbsoluteContext) GetSource() antlr.Token { return s.Source }

func (s *CopyPointerToAbsoluteContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyPointerToAbsoluteContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyPointerToAbsoluteContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyPointerToAbsoluteContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyPointerToAbsoluteContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyPointerToAbsoluteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToAbsoluteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyPointerToAbsoluteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToAbsolute(s)
	}
}

func (s *CopyPointerToAbsoluteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToAbsolute(s)
	}
}

func (p *GrogParser) CopyPointerToAbsolute() (localctx ICopyPointerToAbsoluteContext) {
	localctx = NewCopyPointerToAbsoluteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrogParserRULE_copyPointerToAbsolute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(112)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyPointerToAbsoluteContext).Destination = _m
	}
	{
		p.SetState(113)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyPointerToAbsoluteContext).Source = _m
	}

	return localctx
}

// ICopyAbsoluteToOffsetContext is an interface to support dynamic dispatch.
type ICopyAbsoluteToOffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyAbsoluteToOffsetContext differentiates from other interfaces.
	IsCopyAbsoluteToOffsetContext()
}

type CopyAbsoluteToOffsetContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyAbsoluteToOffsetContext() *CopyAbsoluteToOffsetContext {
	var p = new(CopyAbsoluteToOffsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyAbsoluteToOffset
	return p
}

func (*CopyAbsoluteToOffsetContext) IsCopyAbsoluteToOffsetContext() {}

func NewCopyAbsoluteToOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyAbsoluteToOffsetContext {
	var p = new(CopyAbsoluteToOffsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyAbsoluteToOffset

	return p
}

func (s *CopyAbsoluteToOffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyAbsoluteToOffsetContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyAbsoluteToOffsetContext) GetSource() antlr.Token { return s.Source }

func (s *CopyAbsoluteToOffsetContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyAbsoluteToOffsetContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyAbsoluteToOffsetContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyAbsoluteToOffsetContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyAbsoluteToOffsetContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyAbsoluteToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToOffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyAbsoluteToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToOffset(s)
	}
}

func (s *CopyAbsoluteToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToOffset(s)
	}
}

func (p *GrogParser) CopyAbsoluteToOffset() (localctx ICopyAbsoluteToOffsetContext) {
	localctx = NewCopyAbsoluteToOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrogParserRULE_copyAbsoluteToOffset)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(116)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyAbsoluteToOffsetContext).Destination = _m
	}
	{
		p.SetState(117)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyAbsoluteToOffsetContext).Source = _m
	}

	return localctx
}

// ICopyOffsetToOffsetContext is an interface to support dynamic dispatch.
type ICopyOffsetToOffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyOffsetToOffsetContext differentiates from other interfaces.
	IsCopyOffsetToOffsetContext()
}

type CopyOffsetToOffsetContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyOffsetToOffsetContext() *CopyOffsetToOffsetContext {
	var p = new(CopyOffsetToOffsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyOffsetToOffset
	return p
}

func (*CopyOffsetToOffsetContext) IsCopyOffsetToOffsetContext() {}

func NewCopyOffsetToOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyOffsetToOffsetContext {
	var p = new(CopyOffsetToOffsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyOffsetToOffset

	return p
}

func (s *CopyOffsetToOffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyOffsetToOffsetContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyOffsetToOffsetContext) GetSource() antlr.Token { return s.Source }

func (s *CopyOffsetToOffsetContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyOffsetToOffsetContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyOffsetToOffsetContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyOffsetToOffsetContext) AllOFFSET_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserOFFSET_ADDRESS)
}

func (s *CopyOffsetToOffsetContext) OFFSET_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, i)
}

func (s *CopyOffsetToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToOffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyOffsetToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToOffset(s)
	}
}

func (s *CopyOffsetToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToOffset(s)
	}
}

func (p *GrogParser) CopyOffsetToOffset() (localctx ICopyOffsetToOffsetContext) {
	localctx = NewCopyOffsetToOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrogParserRULE_copyOffsetToOffset)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(120)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyOffsetToOffsetContext).Destination = _m
	}
	{
		p.SetState(121)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyOffsetToOffsetContext).Source = _m
	}

	return localctx
}

// ICopyPointerToOffsetContext is an interface to support dynamic dispatch.
type ICopyPointerToOffsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyPointerToOffsetContext differentiates from other interfaces.
	IsCopyPointerToOffsetContext()
}

type CopyPointerToOffsetContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyPointerToOffsetContext() *CopyPointerToOffsetContext {
	var p = new(CopyPointerToOffsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyPointerToOffset
	return p
}

func (*CopyPointerToOffsetContext) IsCopyPointerToOffsetContext() {}

func NewCopyPointerToOffsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyPointerToOffsetContext {
	var p = new(CopyPointerToOffsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyPointerToOffset

	return p
}

func (s *CopyPointerToOffsetContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyPointerToOffsetContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyPointerToOffsetContext) GetSource() antlr.Token { return s.Source }

func (s *CopyPointerToOffsetContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyPointerToOffsetContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyPointerToOffsetContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyPointerToOffsetContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyPointerToOffsetContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyPointerToOffsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToOffsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyPointerToOffsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToOffset(s)
	}
}

func (s *CopyPointerToOffsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToOffset(s)
	}
}

func (p *GrogParser) CopyPointerToOffset() (localctx ICopyPointerToOffsetContext) {
	localctx = NewCopyPointerToOffsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrogParserRULE_copyPointerToOffset)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(124)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyPointerToOffsetContext).Destination = _m
	}
	{
		p.SetState(125)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyPointerToOffsetContext).Source = _m
	}

	return localctx
}

// ICopyAbsoluteToPointerContext is an interface to support dynamic dispatch.
type ICopyAbsoluteToPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyAbsoluteToPointerContext differentiates from other interfaces.
	IsCopyAbsoluteToPointerContext()
}

type CopyAbsoluteToPointerContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyAbsoluteToPointerContext() *CopyAbsoluteToPointerContext {
	var p = new(CopyAbsoluteToPointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyAbsoluteToPointer
	return p
}

func (*CopyAbsoluteToPointerContext) IsCopyAbsoluteToPointerContext() {}

func NewCopyAbsoluteToPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyAbsoluteToPointerContext {
	var p = new(CopyAbsoluteToPointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyAbsoluteToPointer

	return p
}

func (s *CopyAbsoluteToPointerContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyAbsoluteToPointerContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyAbsoluteToPointerContext) GetSource() antlr.Token { return s.Source }

func (s *CopyAbsoluteToPointerContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyAbsoluteToPointerContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyAbsoluteToPointerContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyAbsoluteToPointerContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyAbsoluteToPointerContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *CopyAbsoluteToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyAbsoluteToPointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyAbsoluteToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyAbsoluteToPointer(s)
	}
}

func (s *CopyAbsoluteToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyAbsoluteToPointer(s)
	}
}

func (p *GrogParser) CopyAbsoluteToPointer() (localctx ICopyAbsoluteToPointerContext) {
	localctx = NewCopyAbsoluteToPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrogParserRULE_copyAbsoluteToPointer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(128)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyAbsoluteToPointerContext).Destination = _m
	}
	{
		p.SetState(129)

		var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

		localctx.(*CopyAbsoluteToPointerContext).Source = _m
	}

	return localctx
}

// ICopyOffsetToPointerContext is an interface to support dynamic dispatch.
type ICopyOffsetToPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyOffsetToPointerContext differentiates from other interfaces.
	IsCopyOffsetToPointerContext()
}

type CopyOffsetToPointerContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyOffsetToPointerContext() *CopyOffsetToPointerContext {
	var p = new(CopyOffsetToPointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyOffsetToPointer
	return p
}

func (*CopyOffsetToPointerContext) IsCopyOffsetToPointerContext() {}

func NewCopyOffsetToPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyOffsetToPointerContext {
	var p = new(CopyOffsetToPointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyOffsetToPointer

	return p
}

func (s *CopyOffsetToPointerContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyOffsetToPointerContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyOffsetToPointerContext) GetSource() antlr.Token { return s.Source }

func (s *CopyOffsetToPointerContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyOffsetToPointerContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyOffsetToPointerContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyOffsetToPointerContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *CopyOffsetToPointerContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *CopyOffsetToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyOffsetToPointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyOffsetToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyOffsetToPointer(s)
	}
}

func (s *CopyOffsetToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyOffsetToPointer(s)
	}
}

func (p *GrogParser) CopyOffsetToPointer() (localctx ICopyOffsetToPointerContext) {
	localctx = NewCopyOffsetToPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrogParserRULE_copyOffsetToPointer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(132)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyOffsetToPointerContext).Destination = _m
	}
	{
		p.SetState(133)

		var _m = p.Match(GrogParserOFFSET_ADDRESS)

		localctx.(*CopyOffsetToPointerContext).Source = _m
	}

	return localctx
}

// ICopyPointerToPointerContext is an interface to support dynamic dispatch.
type ICopyPointerToPointerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsCopyPointerToPointerContext differentiates from other interfaces.
	IsCopyPointerToPointerContext()
}

type CopyPointerToPointerContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyCopyPointerToPointerContext() *CopyPointerToPointerContext {
	var p = new(CopyPointerToPointerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_copyPointerToPointer
	return p
}

func (*CopyPointerToPointerContext) IsCopyPointerToPointerContext() {}

func NewCopyPointerToPointerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyPointerToPointerContext {
	var p = new(CopyPointerToPointerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_copyPointerToPointer

	return p
}

func (s *CopyPointerToPointerContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyPointerToPointerContext) GetDestination() antlr.Token { return s.Destination }

func (s *CopyPointerToPointerContext) GetSource() antlr.Token { return s.Source }

func (s *CopyPointerToPointerContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *CopyPointerToPointerContext) SetSource(v antlr.Token) { s.Source = v }

func (s *CopyPointerToPointerContext) COPY() antlr.TerminalNode {
	return s.GetToken(GrogParserCOPY, 0)
}

func (s *CopyPointerToPointerContext) AllPOINTER_ADDRESS() []antlr.TerminalNode {
	return s.GetTokens(GrogParserPOINTER_ADDRESS)
}

func (s *CopyPointerToPointerContext) POINTER_ADDRESS(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, i)
}

func (s *CopyPointerToPointerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyPointerToPointerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyPointerToPointerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterCopyPointerToPointer(s)
	}
}

func (s *CopyPointerToPointerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitCopyPointerToPointer(s)
	}
}

func (p *GrogParser) CopyPointerToPointer() (localctx ICopyPointerToPointerContext) {
	localctx = NewCopyPointerToPointerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrogParserRULE_copyPointerToPointer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(GrogParserCOPY)
	}
	{
		p.SetState(136)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyPointerToPointerContext).Destination = _m
	}
	{
		p.SetState(137)

		var _m = p.Match(GrogParserPOINTER_ADDRESS)

		localctx.(*CopyPointerToPointerContext).Source = _m
	}

	return localctx
}

// IIncrementContext is an interface to support dynamic dispatch.
type IIncrementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// IsIncrementContext differentiates from other interfaces.
	IsIncrementContext()
}

type IncrementContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
}

func NewEmptyIncrementContext() *IncrementContext {
	var p = new(IncrementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_increment
	return p
}

func (*IncrementContext) IsIncrementContext() {}

func NewIncrementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrementContext {
	var p = new(IncrementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_increment

	return p
}

func (s *IncrementContext) GetParser() antlr.Parser { return s.parser }

func (s *IncrementContext) GetRegister() antlr.Token { return s.Register }

func (s *IncrementContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *IncrementContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(GrogParserINCREMENT, 0)
}

func (s *IncrementContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *IncrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterIncrement(s)
	}
}

func (s *IncrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitIncrement(s)
	}
}

func (p *GrogParser) Increment() (localctx IIncrementContext) {
	localctx = NewIncrementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrogParserRULE_increment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(GrogParserINCREMENT)
	}
	{
		p.SetState(140)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*IncrementContext).Register = _m
	}

	return localctx
}

// IDecrementContext is an interface to support dynamic dispatch.
type IDecrementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRegister returns the Register token.
	GetRegister() antlr.Token

	// SetRegister sets the Register token.
	SetRegister(antlr.Token)

	// IsDecrementContext differentiates from other interfaces.
	IsDecrementContext()
}

type DecrementContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Register antlr.Token
}

func NewEmptyDecrementContext() *DecrementContext {
	var p = new(DecrementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_decrement
	return p
}

func (*DecrementContext) IsDecrementContext() {}

func NewDecrementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecrementContext {
	var p = new(DecrementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_decrement

	return p
}

func (s *DecrementContext) GetParser() antlr.Parser { return s.parser }

func (s *DecrementContext) GetRegister() antlr.Token { return s.Register }

func (s *DecrementContext) SetRegister(v antlr.Token) { s.Register = v }

func (s *DecrementContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(GrogParserDECREMENT, 0)
}

func (s *DecrementContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *DecrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecrementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterDecrement(s)
	}
}

func (s *DecrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitDecrement(s)
	}
}

func (p *GrogParser) Decrement() (localctx IDecrementContext) {
	localctx = NewDecrementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrogParserRULE_decrement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(GrogParserDECREMENT)
	}
	{
		p.SetState(143)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*DecrementContext).Register = _m
	}

	return localctx
}

// IArithmeticOperationContext is an interface to support dynamic dispatch.
type IArithmeticOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsArithmeticOperationContext differentiates from other interfaces.
	IsArithmeticOperationContext()
}

type ArithmeticOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Operator    antlr.Token
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyArithmeticOperationContext() *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_arithmeticOperation
	return p
}

func (*ArithmeticOperationContext) IsArithmeticOperationContext() {}

func NewArithmeticOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticOperationContext {
	var p = new(ArithmeticOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_arithmeticOperation

	return p
}

func (s *ArithmeticOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticOperationContext) GetOperator() antlr.Token { return s.Operator }

func (s *ArithmeticOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *ArithmeticOperationContext) GetSource() antlr.Token { return s.Source }

func (s *ArithmeticOperationContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *ArithmeticOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *ArithmeticOperationContext) SetSource(v antlr.Token) { s.Source = v }

func (s *ArithmeticOperationContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *ArithmeticOperationContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *ArithmeticOperationContext) ADD() antlr.TerminalNode {
	return s.GetToken(GrogParserADD, 0)
}

func (s *ArithmeticOperationContext) SUBTRACT() antlr.TerminalNode {
	return s.GetToken(GrogParserSUBTRACT, 0)
}

func (s *ArithmeticOperationContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(GrogParserMULTIPLY, 0)
}

func (s *ArithmeticOperationContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(GrogParserDIVIDE, 0)
}

func (s *ArithmeticOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterArithmeticOperation(s)
	}
}

func (s *ArithmeticOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitArithmeticOperation(s)
	}
}

func (p *GrogParser) ArithmeticOperation() (localctx IArithmeticOperationContext) {
	localctx = NewArithmeticOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrogParserRULE_arithmeticOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserADD:
		{
			p.SetState(145)

			var _m = p.Match(GrogParserADD)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserSUBTRACT:
		{
			p.SetState(146)

			var _m = p.Match(GrogParserSUBTRACT)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserMULTIPLY:
		{
			p.SetState(147)

			var _m = p.Match(GrogParserMULTIPLY)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	case GrogParserDIVIDE:
		{
			p.SetState(148)

			var _m = p.Match(GrogParserDIVIDE)

			localctx.(*ArithmeticOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(151)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Destination = _m
	}
	{
		p.SetState(152)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*ArithmeticOperationContext).Source = _m
	}

	return localctx
}

// IUnaryBooleanOperationContext is an interface to support dynamic dispatch.
type IUnaryBooleanOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// IsUnaryBooleanOperationContext differentiates from other interfaces.
	IsUnaryBooleanOperationContext()
}

type UnaryBooleanOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Destination antlr.Token
}

func NewEmptyUnaryBooleanOperationContext() *UnaryBooleanOperationContext {
	var p = new(UnaryBooleanOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_unaryBooleanOperation
	return p
}

func (*UnaryBooleanOperationContext) IsUnaryBooleanOperationContext() {}

func NewUnaryBooleanOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryBooleanOperationContext {
	var p = new(UnaryBooleanOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_unaryBooleanOperation

	return p
}

func (s *UnaryBooleanOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryBooleanOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *UnaryBooleanOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *UnaryBooleanOperationContext) NOT() antlr.TerminalNode {
	return s.GetToken(GrogParserNOT, 0)
}

func (s *UnaryBooleanOperationContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *UnaryBooleanOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryBooleanOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryBooleanOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterUnaryBooleanOperation(s)
	}
}

func (s *UnaryBooleanOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitUnaryBooleanOperation(s)
	}
}

func (p *GrogParser) UnaryBooleanOperation() (localctx IUnaryBooleanOperationContext) {
	localctx = NewUnaryBooleanOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrogParserRULE_unaryBooleanOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Match(GrogParserNOT)
	}
	{
		p.SetState(155)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*UnaryBooleanOperationContext).Destination = _m
	}

	return localctx
}

// IBinaryBooleanOperationContext is an interface to support dynamic dispatch.
type IBinaryBooleanOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsBinaryBooleanOperationContext differentiates from other interfaces.
	IsBinaryBooleanOperationContext()
}

type BinaryBooleanOperationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Operator    antlr.Token
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyBinaryBooleanOperationContext() *BinaryBooleanOperationContext {
	var p = new(BinaryBooleanOperationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_binaryBooleanOperation
	return p
}

func (*BinaryBooleanOperationContext) IsBinaryBooleanOperationContext() {}

func NewBinaryBooleanOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryBooleanOperationContext {
	var p = new(BinaryBooleanOperationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_binaryBooleanOperation

	return p
}

func (s *BinaryBooleanOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryBooleanOperationContext) GetOperator() antlr.Token { return s.Operator }

func (s *BinaryBooleanOperationContext) GetDestination() antlr.Token { return s.Destination }

func (s *BinaryBooleanOperationContext) GetSource() antlr.Token { return s.Source }

func (s *BinaryBooleanOperationContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *BinaryBooleanOperationContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *BinaryBooleanOperationContext) SetSource(v antlr.Token) { s.Source = v }

func (s *BinaryBooleanOperationContext) AllREGISTER() []antlr.TerminalNode {
	return s.GetTokens(GrogParserREGISTER)
}

func (s *BinaryBooleanOperationContext) REGISTER(i int) antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, i)
}

func (s *BinaryBooleanOperationContext) AND() antlr.TerminalNode {
	return s.GetToken(GrogParserAND, 0)
}

func (s *BinaryBooleanOperationContext) OR() antlr.TerminalNode {
	return s.GetToken(GrogParserOR, 0)
}

func (s *BinaryBooleanOperationContext) XOR() antlr.TerminalNode {
	return s.GetToken(GrogParserXOR, 0)
}

func (s *BinaryBooleanOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryBooleanOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryBooleanOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterBinaryBooleanOperation(s)
	}
}

func (s *BinaryBooleanOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitBinaryBooleanOperation(s)
	}
}

func (p *GrogParser) BinaryBooleanOperation() (localctx IBinaryBooleanOperationContext) {
	localctx = NewBinaryBooleanOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrogParserRULE_binaryBooleanOperation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserAND:
		{
			p.SetState(157)

			var _m = p.Match(GrogParserAND)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserOR:
		{
			p.SetState(158)

			var _m = p.Match(GrogParserOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	case GrogParserXOR:
		{
			p.SetState(159)

			var _m = p.Match(GrogParserXOR)

			localctx.(*BinaryBooleanOperationContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(162)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Destination = _m
	}
	{
		p.SetState(163)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*BinaryBooleanOperationContext).Source = _m
	}

	return localctx
}

// IJumpContext is an interface to support dynamic dispatch.
type IJumpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the Operator token.
	GetOperator() antlr.Token

	// GetAddress returns the Address token.
	GetAddress() antlr.Token

	// GetOffset returns the Offset token.
	GetOffset() antlr.Token

	// GetPointer returns the Pointer token.
	GetPointer() antlr.Token

	// SetOperator sets the Operator token.
	SetOperator(antlr.Token)

	// SetAddress sets the Address token.
	SetAddress(antlr.Token)

	// SetOffset sets the Offset token.
	SetOffset(antlr.Token)

	// SetPointer sets the Pointer token.
	SetPointer(antlr.Token)

	// IsJumpContext differentiates from other interfaces.
	IsJumpContext()
}

type JumpContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	Operator antlr.Token
	Address  antlr.Token
	Offset   antlr.Token
	Pointer  antlr.Token
}

func NewEmptyJumpContext() *JumpContext {
	var p = new(JumpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_jump
	return p
}

func (*JumpContext) IsJumpContext() {}

func NewJumpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JumpContext {
	var p = new(JumpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_jump

	return p
}

func (s *JumpContext) GetParser() antlr.Parser { return s.parser }

func (s *JumpContext) GetOperator() antlr.Token { return s.Operator }

func (s *JumpContext) GetAddress() antlr.Token { return s.Address }

func (s *JumpContext) GetOffset() antlr.Token { return s.Offset }

func (s *JumpContext) GetPointer() antlr.Token { return s.Pointer }

func (s *JumpContext) SetOperator(v antlr.Token) { s.Operator = v }

func (s *JumpContext) SetAddress(v antlr.Token) { s.Address = v }

func (s *JumpContext) SetOffset(v antlr.Token) { s.Offset = v }

func (s *JumpContext) SetPointer(v antlr.Token) { s.Pointer = v }

func (s *JumpContext) JUMP() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP, 0)
}

func (s *JumpContext) JUMP_IF_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_EQUAL, 0)
}

func (s *JumpContext) JUMP_IF_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_NOT_EQUAL, 0)
}

func (s *JumpContext) JUMP_IF_GREATER() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_GREATER, 0)
}

func (s *JumpContext) JUMP_IF_NOT_GREATER() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_NOT_GREATER, 0)
}

func (s *JumpContext) JUMP_IF_GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_GREATER_OR_EQUAL, 0)
}

func (s *JumpContext) JUMP_IF_NOT_GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL, 0)
}

func (s *JumpContext) JUMP_IF_LESS() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_LESS, 0)
}

func (s *JumpContext) JUMP_IF_NOT_LESS() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_NOT_LESS, 0)
}

func (s *JumpContext) JUMP_IF_LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_LESS_OR_EQUAL, 0)
}

func (s *JumpContext) JUMP_IF_NOT_LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(GrogParserJUMP_IF_NOT_LESS_OR_EQUAL, 0)
}

func (s *JumpContext) ABSOLUTE_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserABSOLUTE_ADDRESS, 0)
}

func (s *JumpContext) OFFSET_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserOFFSET_ADDRESS, 0)
}

func (s *JumpContext) POINTER_ADDRESS() antlr.TerminalNode {
	return s.GetToken(GrogParserPOINTER_ADDRESS, 0)
}

func (s *JumpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JumpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterJump(s)
	}
}

func (s *JumpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitJump(s)
	}
}

func (p *GrogParser) Jump() (localctx IJumpContext) {
	localctx = NewJumpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrogParserRULE_jump)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserJUMP:
		{
			p.SetState(165)

			var _m = p.Match(GrogParserJUMP)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_EQUAL:
		{
			p.SetState(166)

			var _m = p.Match(GrogParserJUMP_IF_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_NOT_EQUAL:
		{
			p.SetState(167)

			var _m = p.Match(GrogParserJUMP_IF_NOT_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_GREATER:
		{
			p.SetState(168)

			var _m = p.Match(GrogParserJUMP_IF_GREATER)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_NOT_GREATER:
		{
			p.SetState(169)

			var _m = p.Match(GrogParserJUMP_IF_NOT_GREATER)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_GREATER_OR_EQUAL:
		{
			p.SetState(170)

			var _m = p.Match(GrogParserJUMP_IF_GREATER_OR_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL:
		{
			p.SetState(171)

			var _m = p.Match(GrogParserJUMP_IF_NOT_GREATER_OR_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_LESS:
		{
			p.SetState(172)

			var _m = p.Match(GrogParserJUMP_IF_LESS)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_NOT_LESS:
		{
			p.SetState(173)

			var _m = p.Match(GrogParserJUMP_IF_NOT_LESS)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_LESS_OR_EQUAL:
		{
			p.SetState(174)

			var _m = p.Match(GrogParserJUMP_IF_LESS_OR_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	case GrogParserJUMP_IF_NOT_LESS_OR_EQUAL:
		{
			p.SetState(175)

			var _m = p.Match(GrogParserJUMP_IF_NOT_LESS_OR_EQUAL)

			localctx.(*JumpContext).Operator = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserABSOLUTE_ADDRESS:
		{
			p.SetState(178)

			var _m = p.Match(GrogParserABSOLUTE_ADDRESS)

			localctx.(*JumpContext).Address = _m
		}

	case GrogParserOFFSET_ADDRESS:
		{
			p.SetState(179)

			var _m = p.Match(GrogParserOFFSET_ADDRESS)

			localctx.(*JumpContext).Offset = _m
		}

	case GrogParserPOINTER_ADDRESS:
		{
			p.SetState(180)

			var _m = p.Match(GrogParserPOINTER_ADDRESS)

			localctx.(*JumpContext).Pointer = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIoContext is an interface to support dynamic dispatch.
type IIoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperation returns the Operation token.
	GetOperation() antlr.Token

	// GetDestination returns the Destination token.
	GetDestination() antlr.Token

	// GetSource returns the Source token.
	GetSource() antlr.Token

	// SetOperation sets the Operation token.
	SetOperation(antlr.Token)

	// SetDestination sets the Destination token.
	SetDestination(antlr.Token)

	// SetSource sets the Source token.
	SetSource(antlr.Token)

	// IsIoContext differentiates from other interfaces.
	IsIoContext()
}

type IoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	Operation   antlr.Token
	Destination antlr.Token
	Source      antlr.Token
}

func NewEmptyIoContext() *IoContext {
	var p = new(IoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_io
	return p
}

func (*IoContext) IsIoContext() {}

func NewIoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IoContext {
	var p = new(IoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_io

	return p
}

func (s *IoContext) GetParser() antlr.Parser { return s.parser }

func (s *IoContext) GetOperation() antlr.Token { return s.Operation }

func (s *IoContext) GetDestination() antlr.Token { return s.Destination }

func (s *IoContext) GetSource() antlr.Token { return s.Source }

func (s *IoContext) SetOperation(v antlr.Token) { s.Operation = v }

func (s *IoContext) SetDestination(v antlr.Token) { s.Destination = v }

func (s *IoContext) SetSource(v antlr.Token) { s.Source = v }

func (s *IoContext) REGISTER() antlr.TerminalNode {
	return s.GetToken(GrogParserREGISTER, 0)
}

func (s *IoContext) DEVICE() antlr.TerminalNode {
	return s.GetToken(GrogParserDEVICE, 0)
}

func (s *IoContext) INPUT() antlr.TerminalNode {
	return s.GetToken(GrogParserINPUT, 0)
}

func (s *IoContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(GrogParserOUTPUT, 0)
}

func (s *IoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterIo(s)
	}
}

func (s *IoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitIo(s)
	}
}

func (p *GrogParser) Io() (localctx IIoContext) {
	localctx = NewIoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrogParserRULE_io)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrogParserINPUT:
		{
			p.SetState(183)

			var _m = p.Match(GrogParserINPUT)

			localctx.(*IoContext).Operation = _m
		}

	case GrogParserOUTPUT:
		{
			p.SetState(184)

			var _m = p.Match(GrogParserOUTPUT)

			localctx.(*IoContext).Operation = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(187)

		var _m = p.Match(GrogParserREGISTER)

		localctx.(*IoContext).Destination = _m
	}
	{
		p.SetState(188)

		var _m = p.Match(GrogParserDEVICE)

		localctx.(*IoContext).Source = _m
	}

	return localctx
}

// IStopContext is an interface to support dynamic dispatch.
type IStopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStopContext differentiates from other interfaces.
	IsStopContext()
}

type StopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStopContext() *StopContext {
	var p = new(StopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_stop
	return p
}

func (*StopContext) IsStopContext() {}

func NewStopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StopContext {
	var p = new(StopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_stop

	return p
}

func (s *StopContext) GetParser() antlr.Parser { return s.parser }

func (s *StopContext) STOP() antlr.TerminalNode {
	return s.GetToken(GrogParserSTOP, 0)
}

func (s *StopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterStop(s)
	}
}

func (s *StopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitStop(s)
	}
}

func (p *GrogParser) Stop() (localctx IStopContext) {
	localctx = NewStopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GrogParserRULE_stop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(GrogParserSTOP)
	}

	return localctx
}

// IWaitContext is an interface to support dynamic dispatch.
type IWaitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWaitContext differentiates from other interfaces.
	IsWaitContext()
}

type WaitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWaitContext() *WaitContext {
	var p = new(WaitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrogParserRULE_wait
	return p
}

func (*WaitContext) IsWaitContext() {}

func NewWaitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaitContext {
	var p = new(WaitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrogParserRULE_wait

	return p
}

func (s *WaitContext) GetParser() antlr.Parser { return s.parser }

func (s *WaitContext) WAIT() antlr.TerminalNode {
	return s.GetToken(GrogParserWAIT, 0)
}

func (s *WaitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WaitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.EnterWait(s)
	}
}

func (s *WaitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrogListener); ok {
		listenerT.ExitWait(s)
	}
}

func (p *GrogParser) Wait() (localctx IWaitContext) {
	localctx = NewWaitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GrogParserRULE_wait)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GrogParserWAIT)
	}

	return localctx
}
