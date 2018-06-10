package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "Input",
			pos:  position{line: 10, col: 1, offset: 126},
			expr: &actionExpr{
				pos: position{line: 10, col: 10, offset: 135},
				run: (*parser).callonInput1,
				expr: &seqExpr{
					pos: position{line: 10, col: 10, offset: 135},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 10, col: 10, offset: 135},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 10, col: 12, offset: 137},
							label: "cl",
							expr: &ruleRefExpr{
								pos:  position{line: 10, col: 16, offset: 141},
								name: "Clauses",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 10, col: 24, offset: 149},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 10, col: 26, offset: 151},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Clauses",
			pos:  position{line: 14, col: 1, offset: 179},
			expr: &actionExpr{
				pos: position{line: 14, col: 12, offset: 190},
				run: (*parser).callonClauses1,
				expr: &seqExpr{
					pos: position{line: 14, col: 12, offset: 190},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 14, col: 12, offset: 190},
							label: "variable",
							expr: &zeroOrOneExpr{
								pos: position{line: 14, col: 21, offset: 199},
								expr: &seqExpr{
									pos: position{line: 14, col: 22, offset: 200},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 14, col: 22, offset: 200},
											name: "Variable",
										},
										&ruleRefExpr{
											pos:  position{line: 14, col: 31, offset: 209},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 14, col: 37, offset: 215},
							label: "constant",
							expr: &zeroOrOneExpr{
								pos: position{line: 14, col: 46, offset: 224},
								expr: &seqExpr{
									pos: position{line: 14, col: 47, offset: 225},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 14, col: 47, offset: 225},
											name: "Constant",
										},
										&ruleRefExpr{
											pos:  position{line: 14, col: 56, offset: 234},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 14, col: 61, offset: 239},
							label: "init",
							expr: &zeroOrOneExpr{
								pos: position{line: 14, col: 66, offset: 244},
								expr: &seqExpr{
									pos: position{line: 14, col: 67, offset: 245},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 14, col: 67, offset: 245},
											name: "Initialization",
										},
										&ruleRefExpr{
											pos:  position{line: 14, col: 82, offset: 260},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 14, col: 87, offset: 265},
							label: "cmd",
							expr: &choiceExpr{
								pos: position{line: 14, col: 93, offset: 271},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 14, col: 93, offset: 271},
										name: "Sequence",
									},
									&ruleRefExpr{
										pos:  position{line: 14, col: 104, offset: 282},
										name: "Command",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 18, col: 1, offset: 356},
			expr: &actionExpr{
				pos: position{line: 18, col: 13, offset: 368},
				run: (*parser).callonVariable1,
				expr: &seqExpr{
					pos: position{line: 18, col: 13, offset: 368},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 18, col: 13, offset: 368},
							val:        "var",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 19, offset: 374},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 18, col: 22, offset: 377},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 18, col: 25, offset: 380},
								name: "Identifier",
							},
						},
						&labeledExpr{
							pos:   position{line: 18, col: 36, offset: 391},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 18, col: 41, offset: 396},
								expr: &seqExpr{
									pos: position{line: 18, col: 43, offset: 398},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 18, col: 43, offset: 398},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 18, col: 45, offset: 400},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 18, col: 49, offset: 404},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 18, col: 51, offset: 406},
											name: "Identifier",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Constant",
			pos:  position{line: 22, col: 1, offset: 473},
			expr: &actionExpr{
				pos: position{line: 22, col: 13, offset: 485},
				run: (*parser).callonConstant1,
				expr: &seqExpr{
					pos: position{line: 22, col: 13, offset: 485},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 22, col: 13, offset: 485},
							val:        "const",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 22, col: 21, offset: 493},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 22, col: 24, offset: 496},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 22, col: 27, offset: 499},
								name: "Identifier",
							},
						},
						&labeledExpr{
							pos:   position{line: 22, col: 38, offset: 510},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 22, col: 43, offset: 515},
								expr: &seqExpr{
									pos: position{line: 22, col: 45, offset: 517},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 22, col: 45, offset: 517},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 22, col: 47, offset: 519},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 22, col: 51, offset: 523},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 22, col: 53, offset: 525},
											name: "Identifier",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Initialization",
			pos:  position{line: 26, col: 1, offset: 592},
			expr: &actionExpr{
				pos: position{line: 26, col: 19, offset: 610},
				run: (*parser).callonInitialization1,
				expr: &seqExpr{
					pos: position{line: 26, col: 19, offset: 610},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 26, col: 19, offset: 610},
							val:        "init",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 26, col: 26, offset: 617},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 26, col: 29, offset: 620},
							label: "sInit",
							expr: &ruleRefExpr{
								pos:  position{line: 26, col: 35, offset: 626},
								name: "SingleInit",
							},
						},
						&labeledExpr{
							pos:   position{line: 26, col: 46, offset: 637},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 26, col: 51, offset: 642},
								expr: &seqExpr{
									pos: position{line: 26, col: 53, offset: 644},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 26, col: 53, offset: 644},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 26, col: 55, offset: 646},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 26, col: 59, offset: 650},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 26, col: 61, offset: 652},
											name: "SingleInit",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "SingleInit",
			pos:  position{line: 30, col: 1, offset: 719},
			expr: &actionExpr{
				pos: position{line: 30, col: 15, offset: 733},
				run: (*parser).callonSingleInit1,
				expr: &seqExpr{
					pos: position{line: 30, col: 15, offset: 733},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 30, col: 15, offset: 733},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 30, col: 19, offset: 737},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 30, col: 30, offset: 748},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 30, col: 32, offset: 750},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 30, col: 36, offset: 754},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 30, col: 38, offset: 756},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 30, col: 44, offset: 762},
								name: "Expression",
							},
						},
					},
				},
			},
		},
		{
			name: "Command",
			pos:  position{line: 34, col: 1, offset: 819},
			expr: &choiceExpr{
				pos: position{line: 34, col: 12, offset: 830},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 34, col: 12, offset: 830},
						name: "While",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 12, offset: 849},
						name: "If",
					},
					&ruleRefExpr{
						pos:  position{line: 36, col: 12, offset: 865},
						name: "Print",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 12, offset: 885},
						name: "Assignment",
					},
				},
			},
		},
		{
			name: "Sequence",
			pos:  position{line: 39, col: 1, offset: 897},
			expr: &actionExpr{
				pos: position{line: 39, col: 13, offset: 909},
				run: (*parser).callonSequence1,
				expr: &seqExpr{
					pos: position{line: 39, col: 13, offset: 909},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 39, col: 13, offset: 909},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 39, col: 19, offset: 915},
								name: "Command",
							},
						},
						&litMatcher{
							pos:        position{line: 39, col: 27, offset: 923},
							val:        ";",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 39, col: 31, offset: 927},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 39, col: 36, offset: 932},
								expr: &seqExpr{
									pos: position{line: 39, col: 38, offset: 934},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 39, col: 38, offset: 934},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 39, col: 40, offset: 936},
											name: "Command",
										},
										&litMatcher{
											pos:        position{line: 39, col: 48, offset: 944},
											val:        ";",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Assignment",
			pos:  position{line: 43, col: 1, offset: 998},
			expr: &actionExpr{
				pos: position{line: 43, col: 15, offset: 1012},
				run: (*parser).callonAssignment1,
				expr: &seqExpr{
					pos: position{line: 43, col: 15, offset: 1012},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 43, col: 15, offset: 1012},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 18, offset: 1015},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 29, offset: 1026},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 43, col: 31, offset: 1028},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 36, offset: 1033},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 38, offset: 1035},
							label: "expr",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 43, offset: 1040},
								name: "Expression",
							},
						},
					},
				},
			},
		},
		{
			name: "While",
			pos:  position{line: 47, col: 1, offset: 1097},
			expr: &actionExpr{
				pos: position{line: 47, col: 10, offset: 1106},
				run: (*parser).callonWhile1,
				expr: &seqExpr{
					pos: position{line: 47, col: 10, offset: 1106},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 47, col: 10, offset: 1106},
							val:        "while",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 18, offset: 1114},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 47, col: 21, offset: 1117},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 25, offset: 1121},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 27, offset: 1123},
							label: "boolExp",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 35, offset: 1131},
								name: "BoolExp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 43, offset: 1139},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 47, col: 45, offset: 1141},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 49, offset: 1145},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 47, col: 52, offset: 1148},
							val:        "do",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 57, offset: 1153},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 59, offset: 1155},
							label: "body",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 64, offset: 1160},
								name: "Block",
							},
						},
					},
				},
			},
		},
		{
			name: "If",
			pos:  position{line: 51, col: 1, offset: 1212},
			expr: &actionExpr{
				pos: position{line: 51, col: 7, offset: 1218},
				run: (*parser).callonIf1,
				expr: &seqExpr{
					pos: position{line: 51, col: 7, offset: 1218},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 51, col: 7, offset: 1218},
							val:        "if",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 12, offset: 1223},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 14, offset: 1225},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 18, offset: 1229},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 20, offset: 1231},
							label: "boolExp",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 28, offset: 1239},
								name: "BoolExp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 36, offset: 1247},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 38, offset: 1249},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 42, offset: 1253},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 44, offset: 1255},
							label: "ifBody",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 51, offset: 1262},
								name: "Block",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 57, offset: 1268},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 59, offset: 1270},
							label: "elseStatement",
							expr: &zeroOrOneExpr{
								pos: position{line: 51, col: 73, offset: 1284},
								expr: &seqExpr{
									pos: position{line: 51, col: 74, offset: 1285},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 51, col: 74, offset: 1285},
											val:        "else",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 51, col: 81, offset: 1292},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 51, col: 83, offset: 1294},
											name: "Block",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Print",
			pos:  position{line: 55, col: 1, offset: 1362},
			expr: &actionExpr{
				pos: position{line: 55, col: 10, offset: 1371},
				run: (*parser).callonPrint1,
				expr: &seqExpr{
					pos: position{line: 55, col: 10, offset: 1371},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 55, col: 10, offset: 1371},
							val:        "print",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 18, offset: 1379},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 20, offset: 1381},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 24, offset: 1385},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 55, col: 26, offset: 1387},
							label: "exp",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 30, offset: 1391},
								name: "Expression",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 41, offset: 1402},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 43, offset: 1404},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Block",
			pos:  position{line: 59, col: 1, offset: 1444},
			expr: &actionExpr{
				pos: position{line: 59, col: 10, offset: 1453},
				run: (*parser).callonBlock1,
				expr: &seqExpr{
					pos: position{line: 59, col: 10, offset: 1453},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 59, col: 10, offset: 1453},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 14, offset: 1457},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 59, col: 16, offset: 1459},
							label: "decl",
							expr: &zeroOrOneExpr{
								pos: position{line: 59, col: 21, offset: 1464},
								expr: &seqExpr{
									pos: position{line: 59, col: 22, offset: 1465},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 59, col: 22, offset: 1465},
											name: "DeclarationSequence",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 42, offset: 1485},
											name: "__",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 47, offset: 1490},
							label: "cmd",
							expr: &choiceExpr{
								pos: position{line: 59, col: 54, offset: 1497},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 59, col: 54, offset: 1497},
										name: "Sequence",
									},
									&ruleRefExpr{
										pos:  position{line: 59, col: 65, offset: 1508},
										name: "Command",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 76, offset: 1519},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 59, col: 78, offset: 1521},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DeclarationSequence",
			pos:  position{line: 63, col: 1, offset: 1567},
			expr: &choiceExpr{
				pos: position{line: 63, col: 24, offset: 1590},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 63, col: 24, offset: 1590},
						name: "Declaration",
					},
					&seqExpr{
						pos: position{line: 63, col: 38, offset: 1604},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 63, col: 38, offset: 1604},
								name: "Declaration",
							},
							&ruleRefExpr{
								pos:  position{line: 63, col: 50, offset: 1616},
								name: "_",
							},
							&litMatcher{
								pos:        position{line: 63, col: 52, offset: 1618},
								val:        ";",
								ignoreCase: false,
							},
							&ruleRefExpr{
								pos:  position{line: 63, col: 56, offset: 1622},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 63, col: 58, offset: 1624},
								name: "Declaration",
							},
						},
					},
				},
			},
		},
		{
			name: "Declaration",
			pos:  position{line: 65, col: 1, offset: 1637},
			expr: &actionExpr{
				pos: position{line: 65, col: 16, offset: 1652},
				run: (*parser).callonDeclaration1,
				expr: &seqExpr{
					pos: position{line: 65, col: 16, offset: 1652},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 65, col: 16, offset: 1652},
							label: "declOp",
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 23, offset: 1659},
								name: "DeclOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 65, col: 30, offset: 1666},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 65, col: 33, offset: 1669},
							label: "initSeq",
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 41, offset: 1677},
								name: "InitSeq",
							},
						},
					},
				},
			},
		},
		{
			name: "InitSeq",
			pos:  position{line: 69, col: 1, offset: 1739},
			expr: &actionExpr{
				pos: position{line: 69, col: 12, offset: 1750},
				run: (*parser).callonInitSeq1,
				expr: &seqExpr{
					pos: position{line: 69, col: 12, offset: 1750},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 69, col: 12, offset: 1750},
							label: "sInit",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 18, offset: 1756},
								name: "SingleInit",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 29, offset: 1767},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 69, col: 34, offset: 1772},
								expr: &seqExpr{
									pos: position{line: 69, col: 36, offset: 1774},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 69, col: 36, offset: 1774},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 69, col: 38, offset: 1776},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 69, col: 42, offset: 1780},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 69, col: 44, offset: 1782},
											name: "SingleInit",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DeclOp",
			pos:  position{line: 73, col: 1, offset: 1849},
			expr: &actionExpr{
				pos: position{line: 73, col: 11, offset: 1859},
				run: (*parser).callonDeclOp1,
				expr: &choiceExpr{
					pos: position{line: 73, col: 12, offset: 1860},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 73, col: 12, offset: 1860},
							val:        "var",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 73, col: 20, offset: 1868},
							val:        "const",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Expression",
			pos:  position{line: 77, col: 1, offset: 1913},
			expr: &choiceExpr{
				pos: position{line: 77, col: 15, offset: 1927},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 77, col: 15, offset: 1927},
						name: "ArithExpr",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 15, offset: 1953},
						name: "BoolExp",
					},
				},
			},
		},
		{
			name: "ArithExpr",
			pos:  position{line: 80, col: 1, offset: 1962},
			expr: &actionExpr{
				pos: position{line: 80, col: 14, offset: 1975},
				run: (*parser).callonArithExpr1,
				expr: &seqExpr{
					pos: position{line: 80, col: 14, offset: 1975},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 80, col: 14, offset: 1975},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 20, offset: 1981},
								name: "Term",
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 25, offset: 1986},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 80, col: 30, offset: 1991},
								expr: &seqExpr{
									pos: position{line: 80, col: 32, offset: 1993},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 80, col: 32, offset: 1993},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 80, col: 34, offset: 1995},
											name: "AddOp",
										},
										&ruleRefExpr{
											pos:  position{line: 80, col: 40, offset: 2001},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 80, col: 42, offset: 2003},
											name: "Term",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Term",
			pos:  position{line: 84, col: 1, offset: 2059},
			expr: &actionExpr{
				pos: position{line: 84, col: 9, offset: 2067},
				run: (*parser).callonTerm1,
				expr: &seqExpr{
					pos: position{line: 84, col: 9, offset: 2067},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 84, col: 9, offset: 2067},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 84, col: 15, offset: 2073},
								name: "Factor",
							},
						},
						&labeledExpr{
							pos:   position{line: 84, col: 22, offset: 2080},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 84, col: 27, offset: 2085},
								expr: &seqExpr{
									pos: position{line: 84, col: 29, offset: 2087},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 84, col: 29, offset: 2087},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 84, col: 31, offset: 2089},
											name: "MulOp",
										},
										&ruleRefExpr{
											pos:  position{line: 84, col: 37, offset: 2095},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 84, col: 39, offset: 2097},
											name: "Factor",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Factor",
			pos:  position{line: 88, col: 1, offset: 2155},
			expr: &choiceExpr{
				pos: position{line: 88, col: 11, offset: 2165},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 88, col: 11, offset: 2165},
						run: (*parser).callonFactor2,
						expr: &seqExpr{
							pos: position{line: 88, col: 11, offset: 2165},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 88, col: 11, offset: 2165},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 88, col: 15, offset: 2169},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 88, col: 20, offset: 2174},
										name: "ArithExpr",
									},
								},
								&litMatcher{
									pos:        position{line: 88, col: 30, offset: 2184},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 90, col: 5, offset: 2215},
						run: (*parser).callonFactor8,
						expr: &labeledExpr{
							pos:   position{line: 90, col: 5, offset: 2215},
							label: "integer",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 13, offset: 2223},
								name: "Integer",
							},
						},
					},
					&actionExpr{
						pos: position{line: 92, col: 5, offset: 2261},
						run: (*parser).callonFactor11,
						expr: &labeledExpr{
							pos:   position{line: 92, col: 5, offset: 2261},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 9, offset: 2265},
								name: "Identifier",
							},
						},
					},
				},
			},
		},
		{
			name: "AddOp",
			pos:  position{line: 96, col: 1, offset: 2300},
			expr: &actionExpr{
				pos: position{line: 96, col: 10, offset: 2309},
				run: (*parser).callonAddOp1,
				expr: &choiceExpr{
					pos: position{line: 96, col: 11, offset: 2310},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 96, col: 11, offset: 2310},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 96, col: 17, offset: 2316},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MulOp",
			pos:  position{line: 100, col: 1, offset: 2358},
			expr: &actionExpr{
				pos: position{line: 100, col: 10, offset: 2367},
				run: (*parser).callonMulOp1,
				expr: &choiceExpr{
					pos: position{line: 100, col: 12, offset: 2369},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 100, col: 12, offset: 2369},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 100, col: 18, offset: 2375},
							val:        "/",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoolExp",
			pos:  position{line: 104, col: 1, offset: 2417},
			expr: &choiceExpr{
				pos: position{line: 104, col: 12, offset: 2428},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 104, col: 12, offset: 2428},
						run: (*parser).callonBoolExp2,
						expr: &labeledExpr{
							pos:   position{line: 104, col: 12, offset: 2428},
							label: "keyword",
							expr: &ruleRefExpr{
								pos:  position{line: 104, col: 21, offset: 2437},
								name: "KeywordBoolExp",
							},
						},
					},
					&actionExpr{
						pos: position{line: 107, col: 3, offset: 2502},
						run: (*parser).callonBoolExp5,
						expr: &seqExpr{
							pos: position{line: 107, col: 3, offset: 2502},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 107, col: 3, offset: 2502},
									label: "unBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 107, col: 13, offset: 2512},
										name: "UnaryBoolOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 107, col: 25, offset: 2524},
									label: "boolExp",
									expr: &ruleRefExpr{
										pos:  position{line: 107, col: 34, offset: 2533},
										name: "BoolExp",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 110, col: 3, offset: 2599},
						run: (*parser).callonBoolExp11,
						expr: &seqExpr{
							pos: position{line: 110, col: 3, offset: 2599},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 110, col: 3, offset: 2599},
									label: "id1",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 8, offset: 2604},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 19, offset: 2615},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 110, col: 21, offset: 2617},
									label: "binBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 32, offset: 2628},
										name: "BinaryBoolOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 45, offset: 2641},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 110, col: 47, offset: 2643},
									label: "id2",
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 52, offset: 2648},
										name: "Identifier",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 113, col: 3, offset: 2721},
						run: (*parser).callonBoolExp21,
						expr: &seqExpr{
							pos: position{line: 113, col: 3, offset: 2721},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 113, col: 3, offset: 2721},
									label: "id1",
									expr: &ruleRefExpr{
										pos:  position{line: 113, col: 8, offset: 2726},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 113, col: 19, offset: 2737},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 113, col: 21, offset: 2739},
									label: "binBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 113, col: 32, offset: 2750},
										name: "BinaryBoolOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 113, col: 45, offset: 2763},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 113, col: 47, offset: 2765},
									label: "id2",
									expr: &ruleRefExpr{
										pos:  position{line: 113, col: 51, offset: 2769},
										name: "ArithExpr",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "KeywordBoolExp",
			pos:  position{line: 117, col: 1, offset: 2840},
			expr: &actionExpr{
				pos: position{line: 117, col: 19, offset: 2858},
				run: (*parser).callonKeywordBoolExp1,
				expr: &choiceExpr{
					pos: position{line: 117, col: 20, offset: 2859},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 117, col: 20, offset: 2859},
							val:        "true",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 117, col: 29, offset: 2868},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "UnaryBoolOp",
			pos:  position{line: 121, col: 1, offset: 2913},
			expr: &actionExpr{
				pos: position{line: 121, col: 16, offset: 2928},
				run: (*parser).callonUnaryBoolOp1,
				expr: &litMatcher{
					pos:        position{line: 121, col: 16, offset: 2928},
					val:        "~",
					ignoreCase: false,
				},
			},
		},
		{
			name: "BinaryBoolOp",
			pos:  position{line: 125, col: 1, offset: 2967},
			expr: &actionExpr{
				pos: position{line: 125, col: 17, offset: 2983},
				run: (*parser).callonBinaryBoolOp1,
				expr: &choiceExpr{
					pos: position{line: 125, col: 19, offset: 2985},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 125, col: 19, offset: 2985},
							val:        "==",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 26, offset: 2992},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 33, offset: 2999},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 40, offset: 3006},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 46, offset: 3012},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 52, offset: 3018},
							val:        "/\\",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 60, offset: 3026},
							val:        "\\/",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 129, col: 1, offset: 3069},
			expr: &actionExpr{
				pos: position{line: 129, col: 15, offset: 3083},
				run: (*parser).callonIdentifier1,
				expr: &seqExpr{
					pos: position{line: 129, col: 15, offset: 3083},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 129, col: 15, offset: 3083},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 129, col: 23, offset: 3091},
							expr: &charClassMatcher{
								pos:        position{line: 129, col: 23, offset: 3091},
								val:        "[0-9a-zA-Z]",
								ranges:     []rune{'0', '9', 'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "Integer",
			pos:  position{line: 133, col: 1, offset: 3139},
			expr: &actionExpr{
				pos: position{line: 133, col: 12, offset: 3150},
				run: (*parser).callonInteger1,
				expr: &seqExpr{
					pos: position{line: 133, col: 12, offset: 3150},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 133, col: 12, offset: 3150},
							expr: &litMatcher{
								pos:        position{line: 133, col: 12, offset: 3150},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 133, col: 17, offset: 3155},
							expr: &charClassMatcher{
								pos:        position{line: 133, col: 17, offset: 3155},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 137, col: 1, offset: 3198},
			expr: &zeroOrMoreExpr{
				pos: position{line: 137, col: 19, offset: 3216},
				expr: &charClassMatcher{
					pos:        position{line: 137, col: 19, offset: 3216},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "__",
			displayName: "\"obligatory_whitespace\"",
			pos:         position{line: 138, col: 1, offset: 3227},
			expr: &oneOrMoreExpr{
				pos: position{line: 138, col: 31, offset: 3257},
				expr: &charClassMatcher{
					pos:        position{line: 138, col: 31, offset: 3257},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 140, col: 1, offset: 3269},
			expr: &notExpr{
				pos: position{line: 140, col: 8, offset: 3276},
				expr: &anyMatcher{
					line: 140, col: 9, offset: 3277,
				},
			},
		},
	},
}

func (c *current) onInput1(cl interface{}) (interface{}, error) {
	return cl, nil
}

func (p *parser) callonInput1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInput1(stack["cl"])
}

func (c *current) onClauses1(variable, constant, init, cmd interface{}) (interface{}, error) {
	return evalClauses(variable, constant, init, cmd), nil
}

func (p *parser) callonClauses1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onClauses1(stack["variable"], stack["constant"], stack["init"], stack["cmd"])
}

func (c *current) onVariable1(id, rest interface{}) (interface{}, error) {
	return evalClauseDeclaration(id, rest), nil
}

func (p *parser) callonVariable1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onVariable1(stack["id"], stack["rest"])
}

func (c *current) onConstant1(id, rest interface{}) (interface{}, error) {
	return evalClauseDeclaration(id, rest), nil
}

func (p *parser) callonConstant1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onConstant1(stack["id"], stack["rest"])
}

func (c *current) onInitialization1(sInit, rest interface{}) (interface{}, error) {
	return evalInitialization(sInit, rest), nil
}

func (p *parser) callonInitialization1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInitialization1(stack["sInit"], stack["rest"])
}

func (c *current) onSingleInit1(id, expr interface{}) (interface{}, error) {
	return evalSingleInit(id, expr), nil
}

func (p *parser) callonSingleInit1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSingleInit1(stack["id"], stack["expr"])
}

func (c *current) onSequence1(first, rest interface{}) (interface{}, error) {
	return evalSequence(first, rest), nil
}

func (p *parser) callonSequence1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSequence1(stack["first"], stack["rest"])
}

func (c *current) onAssignment1(id, expr interface{}) (interface{}, error) {
	return evalAssignment(id, expr), nil
}

func (p *parser) callonAssignment1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAssignment1(stack["id"], stack["expr"])
}

func (c *current) onWhile1(boolExp, body interface{}) (interface{}, error) {
	return evalWhile(boolExp, body), nil
}

func (p *parser) callonWhile1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWhile1(stack["boolExp"], stack["body"])
}

func (c *current) onIf1(boolExp, ifBody, elseStatement interface{}) (interface{}, error) {
	return evalIf(boolExp, ifBody, elseStatement), nil
}

func (p *parser) callonIf1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIf1(stack["boolExp"], stack["ifBody"], stack["elseStatement"])
}

func (c *current) onPrint1(exp interface{}) (interface{}, error) {
	return evalPrint(exp), nil
}

func (p *parser) callonPrint1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPrint1(stack["exp"])
}

func (c *current) onBlock1(decl, cmd interface{}) (interface{}, error) {
	return evalBlock(decl, cmd), nil
}

func (p *parser) callonBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlock1(stack["decl"], stack["cmd"])
}

func (c *current) onDeclaration1(declOp, initSeq interface{}) (interface{}, error) {
	return evalDeclaration(declOp, initSeq), nil
}

func (p *parser) callonDeclaration1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDeclaration1(stack["declOp"], stack["initSeq"])
}

func (c *current) onInitSeq1(sInit, rest interface{}) (interface{}, error) {
	return evalInitialization(sInit, rest), nil
}

func (p *parser) callonInitSeq1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInitSeq1(stack["sInit"], stack["rest"])
}

func (c *current) onDeclOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonDeclOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDeclOp1()
}

func (c *current) onArithExpr1(first, rest interface{}) (interface{}, error) {
	return evalArithExpr(first, rest), nil
}

func (p *parser) callonArithExpr1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onArithExpr1(stack["first"], stack["rest"])
}

func (c *current) onTerm1(first, rest interface{}) (interface{}, error) {
	return evalArithExpr(first, rest), nil
}

func (p *parser) callonTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTerm1(stack["first"], stack["rest"])
}

func (c *current) onFactor2(expr interface{}) (interface{}, error) {
	return expr, nil
}

func (p *parser) callonFactor2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFactor2(stack["expr"])
}

func (c *current) onFactor8(integer interface{}) (interface{}, error) {
	return integer, nil
}

func (p *parser) callonFactor8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFactor8(stack["integer"])
}

func (c *current) onFactor11(id interface{}) (interface{}, error) {
	return id, nil
}

func (p *parser) callonFactor11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFactor11(stack["id"])
}

func (c *current) onAddOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonAddOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAddOp1()
}

func (c *current) onMulOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonMulOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMulOp1()
}

func (c *current) onBoolExp2(keyword interface{}) (interface{}, error) {
	return evalKeywordBoolExp(keyword), nil
}

func (p *parser) callonBoolExp2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolExp2(stack["keyword"])
}

func (c *current) onBoolExp5(unBoolOp, boolExp interface{}) (interface{}, error) {
	return evalUnaryBoolExp(unBoolOp, boolExp), nil
}

func (p *parser) callonBoolExp5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolExp5(stack["unBoolOp"], stack["boolExp"])
}

func (c *current) onBoolExp11(id1, binBoolOp, id2 interface{}) (interface{}, error) {
	return evalBinaryLogicExp(binBoolOp, id1, id2), nil
}

func (p *parser) callonBoolExp11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolExp11(stack["id1"], stack["binBoolOp"], stack["id2"])
}

func (c *current) onBoolExp21(id1, binBoolOp, id2 interface{}) (interface{}, error) {
	return evalBinaryArithExp(binBoolOp, id1, id2), nil
}

func (p *parser) callonBoolExp21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoolExp21(stack["id1"], stack["binBoolOp"], stack["id2"])
}

func (c *current) onKeywordBoolExp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonKeywordBoolExp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onKeywordBoolExp1()
}

func (c *current) onUnaryBoolOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonUnaryBoolOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnaryBoolOp1()
}

func (c *current) onBinaryBoolOp1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonBinaryBoolOp1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBinaryBoolOp1()
}

func (c *current) onIdentifier1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonIdentifier1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdentifier1()
}

func (c *current) onInteger1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonInteger1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInteger1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// AllowInvalidUTF8 creates an Option to allow invalid UTF-8 bytes.
// Every invalid UTF-8 byte is treated as a utf8.RuneError (U+FFFD)
// by character class matchers and is matched by the any matcher.
// The returned matched value, c.text and c.offset are NOT affected.
//
// The default is false.
func AllowInvalidUTF8(b bool) Option {
	return func(p *parser) Option {
		old := p.allowInvalidUTF8
		p.allowInvalidUTF8 = b
		return AllowInvalidUTF8(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type stateCodeExpr struct {
	pos position
	run func(*parser) error
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	allowInvalidUTF8 bool

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		if !p.allowInvalidUTF8 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		if len(p.emptyState) > 0 {
			p.emptyState = make(storeDict)
		}
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn == utf8.RuneError && p.pt.w == 0 {
		// EOF - see utf8.DecodeRune
		p.failAt(false, p.pt.position, ".")
		return nil, false
	}
	start := p.pt
	p.read()
	p.failAt(true, start.position, ".")
	return p.sliceFrom(start), true
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()

		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	state := p.cloneState()
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restoreState(state)
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}
