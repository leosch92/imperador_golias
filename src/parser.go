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
			pos:  position{line: 15, col: 1, offset: 180},
			expr: &actionExpr{
				pos: position{line: 15, col: 12, offset: 191},
				run: (*parser).callonClauses1,
				expr: &seqExpr{
					pos: position{line: 15, col: 12, offset: 191},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 15, col: 12, offset: 191},
							label: "variable",
							expr: &zeroOrOneExpr{
								pos: position{line: 15, col: 21, offset: 200},
								expr: &seqExpr{
									pos: position{line: 15, col: 22, offset: 201},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 15, col: 22, offset: 201},
											name: "Variable",
										},
										&ruleRefExpr{
											pos:  position{line: 15, col: 31, offset: 210},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 15, col: 33, offset: 212},
											val:        ";",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 15, col: 37, offset: 216},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 15, col: 43, offset: 222},
							label: "constant",
							expr: &zeroOrOneExpr{
								pos: position{line: 15, col: 52, offset: 231},
								expr: &seqExpr{
									pos: position{line: 15, col: 53, offset: 232},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 15, col: 53, offset: 232},
											name: "Constant",
										},
										&ruleRefExpr{
											pos:  position{line: 15, col: 62, offset: 241},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 15, col: 64, offset: 243},
											val:        ";",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 15, col: 68, offset: 247},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 12, offset: 263},
							label: "init",
							expr: &zeroOrOneExpr{
								pos: position{line: 16, col: 17, offset: 268},
								expr: &seqExpr{
									pos: position{line: 16, col: 18, offset: 269},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 16, col: 18, offset: 269},
											name: "Initialization",
										},
										&ruleRefExpr{
											pos:  position{line: 16, col: 33, offset: 284},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 16, col: 35, offset: 286},
											val:        ";",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 42, offset: 293},
							label: "procs",
							expr: &oneOrMoreExpr{
								pos: position{line: 16, col: 48, offset: 299},
								expr: &seqExpr{
									pos: position{line: 16, col: 50, offset: 301},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 16, col: 50, offset: 301},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 16, col: 52, offset: 303},
											name: "Procedure",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 65, offset: 316},
							label: "calls",
							expr: &zeroOrMoreExpr{
								pos: position{line: 16, col: 71, offset: 322},
								expr: &seqExpr{
									pos: position{line: 16, col: 73, offset: 324},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 16, col: 73, offset: 324},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 16, col: 75, offset: 326},
											name: "Call",
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
			name: "Procedure",
			pos:  position{line: 20, col: 1, offset: 412},
			expr: &actionExpr{
				pos: position{line: 20, col: 14, offset: 425},
				run: (*parser).callonProcedure1,
				expr: &seqExpr{
					pos: position{line: 20, col: 14, offset: 425},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 20, col: 14, offset: 425},
							val:        "proc",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 20, col: 21, offset: 432},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 20, col: 24, offset: 435},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 20, col: 27, offset: 438},
								name: "Identifier",
							},
						},
						&litMatcher{
							pos:        position{line: 20, col: 38, offset: 449},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 20, col: 42, offset: 453},
							label: "formals",
							expr: &ruleRefExpr{
								pos:  position{line: 20, col: 51, offset: 462},
								name: "Formals",
							},
						},
						&litMatcher{
							pos:        position{line: 20, col: 59, offset: 470},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 20, col: 63, offset: 474},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 20, col: 65, offset: 476},
							label: "blk",
							expr: &ruleRefExpr{
								pos:  position{line: 20, col: 70, offset: 481},
								name: "Block",
							},
						},
					},
				},
			},
		},
		{
			name: "Formals",
			pos:  position{line: 24, col: 1, offset: 544},
			expr: &choiceExpr{
				pos: position{line: 24, col: 12, offset: 555},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 24, col: 12, offset: 555},
						run: (*parser).callonFormals2,
						expr: &seqExpr{
							pos: position{line: 24, col: 12, offset: 555},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 24, col: 12, offset: 555},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 24, col: 18, offset: 561},
										name: "Identifier",
									},
								},
								&labeledExpr{
									pos:   position{line: 24, col: 29, offset: 572},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 24, col: 34, offset: 577},
										expr: &seqExpr{
											pos: position{line: 24, col: 35, offset: 578},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 24, col: 35, offset: 578},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 24, col: 39, offset: 582},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 24, col: 41, offset: 584},
													name: "Identifier",
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 26, col: 5, offset: 650},
						run: (*parser).callonFormals12,
						expr: &litMatcher{
							pos:        position{line: 26, col: 5, offset: 650},
							val:        "",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Call",
			pos:  position{line: 30, col: 1, offset: 678},
			expr: &actionExpr{
				pos: position{line: 30, col: 9, offset: 686},
				run: (*parser).callonCall1,
				expr: &seqExpr{
					pos: position{line: 30, col: 9, offset: 686},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 30, col: 9, offset: 686},
							val:        "exec",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 30, col: 16, offset: 693},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 30, col: 19, offset: 696},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 30, col: 22, offset: 699},
								name: "Identifier",
							},
						},
						&litMatcher{
							pos:        position{line: 30, col: 33, offset: 710},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 30, col: 37, offset: 714},
							label: "actuals",
							expr: &ruleRefExpr{
								pos:  position{line: 30, col: 46, offset: 723},
								name: "Actuals",
							},
						},
						&litMatcher{
							pos:        position{line: 30, col: 54, offset: 731},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Actuals",
			pos:  position{line: 34, col: 1, offset: 783},
			expr: &choiceExpr{
				pos: position{line: 34, col: 12, offset: 794},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 34, col: 12, offset: 794},
						run: (*parser).callonActuals2,
						expr: &seqExpr{
							pos: position{line: 34, col: 12, offset: 794},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 34, col: 12, offset: 794},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 34, col: 18, offset: 800},
										name: "Expression",
									},
								},
								&labeledExpr{
									pos:   position{line: 34, col: 29, offset: 811},
									label: "rest",
									expr: &zeroOrMoreExpr{
										pos: position{line: 34, col: 34, offset: 816},
										expr: &seqExpr{
											pos: position{line: 34, col: 36, offset: 818},
											exprs: []interface{}{
												&litMatcher{
													pos:        position{line: 34, col: 36, offset: 818},
													val:        ",",
													ignoreCase: false,
												},
												&ruleRefExpr{
													pos:  position{line: 34, col: 40, offset: 822},
													name: "_",
												},
												&ruleRefExpr{
													pos:  position{line: 34, col: 42, offset: 824},
													name: "Expression",
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 36, col: 5, offset: 890},
						run: (*parser).callonActuals12,
						expr: &litMatcher{
							pos:        position{line: 36, col: 5, offset: 890},
							val:        "",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Variable",
			pos:  position{line: 40, col: 1, offset: 918},
			expr: &actionExpr{
				pos: position{line: 40, col: 13, offset: 930},
				run: (*parser).callonVariable1,
				expr: &seqExpr{
					pos: position{line: 40, col: 13, offset: 930},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 40, col: 13, offset: 930},
							val:        "var",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 40, col: 19, offset: 936},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 40, col: 22, offset: 939},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 40, col: 25, offset: 942},
								name: "Identifier",
							},
						},
						&labeledExpr{
							pos:   position{line: 40, col: 36, offset: 953},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 40, col: 41, offset: 958},
								expr: &seqExpr{
									pos: position{line: 40, col: 43, offset: 960},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 40, col: 43, offset: 960},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 40, col: 45, offset: 962},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 40, col: 49, offset: 966},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 40, col: 51, offset: 968},
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
			pos:  position{line: 44, col: 1, offset: 1035},
			expr: &actionExpr{
				pos: position{line: 44, col: 13, offset: 1047},
				run: (*parser).callonConstant1,
				expr: &seqExpr{
					pos: position{line: 44, col: 13, offset: 1047},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 44, col: 13, offset: 1047},
							val:        "const",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 44, col: 21, offset: 1055},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 44, col: 24, offset: 1058},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 44, col: 27, offset: 1061},
								name: "Identifier",
							},
						},
						&labeledExpr{
							pos:   position{line: 44, col: 38, offset: 1072},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 44, col: 43, offset: 1077},
								expr: &seqExpr{
									pos: position{line: 44, col: 45, offset: 1079},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 44, col: 45, offset: 1079},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 44, col: 47, offset: 1081},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 44, col: 51, offset: 1085},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 44, col: 53, offset: 1087},
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
			pos:  position{line: 48, col: 1, offset: 1154},
			expr: &actionExpr{
				pos: position{line: 48, col: 19, offset: 1172},
				run: (*parser).callonInitialization1,
				expr: &seqExpr{
					pos: position{line: 48, col: 19, offset: 1172},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 48, col: 19, offset: 1172},
							val:        "init",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 48, col: 26, offset: 1179},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 48, col: 29, offset: 1182},
							label: "sInit",
							expr: &ruleRefExpr{
								pos:  position{line: 48, col: 35, offset: 1188},
								name: "SingleInit",
							},
						},
						&labeledExpr{
							pos:   position{line: 48, col: 46, offset: 1199},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 48, col: 51, offset: 1204},
								expr: &seqExpr{
									pos: position{line: 48, col: 53, offset: 1206},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 48, col: 53, offset: 1206},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 48, col: 55, offset: 1208},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 48, col: 59, offset: 1212},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 48, col: 61, offset: 1214},
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
			pos:  position{line: 52, col: 1, offset: 1281},
			expr: &actionExpr{
				pos: position{line: 52, col: 15, offset: 1295},
				run: (*parser).callonSingleInit1,
				expr: &seqExpr{
					pos: position{line: 52, col: 15, offset: 1295},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 52, col: 15, offset: 1295},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 19, offset: 1299},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 30, offset: 1310},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 52, col: 32, offset: 1312},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 52, col: 36, offset: 1316},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 52, col: 38, offset: 1318},
							label: "expr",
							expr: &choiceExpr{
								pos: position{line: 52, col: 44, offset: 1324},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 52, col: 44, offset: 1324},
										name: "Expression",
									},
									&ruleRefExpr{
										pos:  position{line: 52, col: 57, offset: 1337},
										name: "Text",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Command",
			pos:  position{line: 56, col: 1, offset: 1389},
			expr: &choiceExpr{
				pos: position{line: 56, col: 12, offset: 1400},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 56, col: 12, offset: 1400},
						name: "While",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 12, offset: 1419},
						name: "If",
					},
					&ruleRefExpr{
						pos:  position{line: 58, col: 12, offset: 1435},
						name: "Print",
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 12, offset: 1454},
						name: "Assignment",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 12, offset: 1478},
						name: "Call",
					},
				},
			},
		},
		{
			name: "Sequence",
			pos:  position{line: 62, col: 1, offset: 1484},
			expr: &actionExpr{
				pos: position{line: 62, col: 13, offset: 1496},
				run: (*parser).callonSequence1,
				expr: &seqExpr{
					pos: position{line: 62, col: 13, offset: 1496},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 62, col: 13, offset: 1496},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 19, offset: 1502},
								name: "Command",
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 27, offset: 1510},
							label: "rest",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 32, offset: 1515},
								expr: &seqExpr{
									pos: position{line: 62, col: 34, offset: 1517},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 62, col: 34, offset: 1517},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 62, col: 36, offset: 1519},
											val:        ";",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 62, col: 40, offset: 1523},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 62, col: 42, offset: 1525},
											name: "Command",
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
			pos:  position{line: 66, col: 1, offset: 1583},
			expr: &actionExpr{
				pos: position{line: 66, col: 15, offset: 1597},
				run: (*parser).callonAssignment1,
				expr: &seqExpr{
					pos: position{line: 66, col: 15, offset: 1597},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 66, col: 15, offset: 1597},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 18, offset: 1600},
								name: "Identifier",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 29, offset: 1611},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 31, offset: 1613},
							val:        ":=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 36, offset: 1618},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 66, col: 38, offset: 1620},
							label: "expr",
							expr: &choiceExpr{
								pos: position{line: 66, col: 44, offset: 1626},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 66, col: 44, offset: 1626},
										name: "Expression",
									},
									&ruleRefExpr{
										pos:  position{line: 66, col: 57, offset: 1639},
										name: "Text",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "While",
			pos:  position{line: 70, col: 1, offset: 1691},
			expr: &actionExpr{
				pos: position{line: 70, col: 10, offset: 1700},
				run: (*parser).callonWhile1,
				expr: &seqExpr{
					pos: position{line: 70, col: 10, offset: 1700},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 70, col: 10, offset: 1700},
							val:        "while",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 18, offset: 1708},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 70, col: 21, offset: 1711},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 25, offset: 1715},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 70, col: 27, offset: 1717},
							label: "boolExp",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 35, offset: 1725},
								name: "BoolExp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 43, offset: 1733},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 70, col: 45, offset: 1735},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 49, offset: 1739},
							name: "__",
						},
						&litMatcher{
							pos:        position{line: 70, col: 52, offset: 1742},
							val:        "do",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 70, col: 57, offset: 1747},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 70, col: 59, offset: 1749},
							label: "body",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 64, offset: 1754},
								name: "Block",
							},
						},
					},
				},
			},
		},
		{
			name: "If",
			pos:  position{line: 74, col: 1, offset: 1806},
			expr: &actionExpr{
				pos: position{line: 74, col: 7, offset: 1812},
				run: (*parser).callonIf1,
				expr: &seqExpr{
					pos: position{line: 74, col: 7, offset: 1812},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 74, col: 7, offset: 1812},
							val:        "if",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 12, offset: 1817},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 74, col: 14, offset: 1819},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 18, offset: 1823},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 74, col: 20, offset: 1825},
							label: "boolExp",
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 28, offset: 1833},
								name: "BoolExp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 36, offset: 1841},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 74, col: 38, offset: 1843},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 42, offset: 1847},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 74, col: 44, offset: 1849},
							label: "ifBody",
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 51, offset: 1856},
								name: "Block",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 57, offset: 1862},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 74, col: 59, offset: 1864},
							label: "elseStatement",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 73, offset: 1878},
								expr: &seqExpr{
									pos: position{line: 74, col: 74, offset: 1879},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 74, col: 74, offset: 1879},
											val:        "else",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 74, col: 81, offset: 1886},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 74, col: 83, offset: 1888},
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
			pos:  position{line: 78, col: 1, offset: 1956},
			expr: &actionExpr{
				pos: position{line: 78, col: 10, offset: 1965},
				run: (*parser).callonPrint1,
				expr: &seqExpr{
					pos: position{line: 78, col: 10, offset: 1965},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 78, col: 10, offset: 1965},
							val:        "print",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 18, offset: 1973},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 78, col: 20, offset: 1975},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 24, offset: 1979},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 78, col: 26, offset: 1981},
							label: "exp",
							expr: &choiceExpr{
								pos: position{line: 78, col: 31, offset: 1986},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 78, col: 31, offset: 1986},
										name: "Expression",
									},
									&ruleRefExpr{
										pos:  position{line: 78, col: 44, offset: 1999},
										name: "Text",
									},
									&ruleRefExpr{
										pos:  position{line: 78, col: 51, offset: 2006},
										name: "Real",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 57, offset: 2012},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 78, col: 59, offset: 2014},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Block",
			pos:  position{line: 82, col: 1, offset: 2054},
			expr: &actionExpr{
				pos: position{line: 82, col: 10, offset: 2063},
				run: (*parser).callonBlock1,
				expr: &seqExpr{
					pos: position{line: 82, col: 10, offset: 2063},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 82, col: 10, offset: 2063},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 14, offset: 2067},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 82, col: 16, offset: 2069},
							label: "decl",
							expr: &zeroOrOneExpr{
								pos: position{line: 82, col: 21, offset: 2074},
								expr: &seqExpr{
									pos: position{line: 82, col: 22, offset: 2075},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 82, col: 22, offset: 2075},
											name: "DeclarationSequence",
										},
										&ruleRefExpr{
											pos:  position{line: 82, col: 42, offset: 2095},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 82, col: 44, offset: 2097},
											val:        ";",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 82, col: 48, offset: 2101},
											name: "_",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 82, col: 53, offset: 2106},
							label: "cmd",
							expr: &choiceExpr{
								pos: position{line: 82, col: 60, offset: 2113},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 82, col: 60, offset: 2113},
										name: "Sequence",
									},
									&ruleRefExpr{
										pos:  position{line: 82, col: 71, offset: 2124},
										name: "Command",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 82, offset: 2135},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 82, col: 84, offset: 2137},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DeclarationSequence",
			pos:  position{line: 86, col: 1, offset: 2183},
			expr: &choiceExpr{
				pos: position{line: 86, col: 24, offset: 2206},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 86, col: 24, offset: 2206},
						run: (*parser).callonDeclarationSequence2,
						expr: &seqExpr{
							pos: position{line: 86, col: 24, offset: 2206},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 86, col: 24, offset: 2206},
									label: "first",
									expr: &ruleRefExpr{
										pos:  position{line: 86, col: 31, offset: 2213},
										name: "Declaration",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 86, col: 43, offset: 2225},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 86, col: 45, offset: 2227},
									val:        ";",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 86, col: 49, offset: 2231},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 86, col: 51, offset: 2233},
									label: "rest",
									expr: &choiceExpr{
										pos: position{line: 86, col: 58, offset: 2240},
										alternatives: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 86, col: 58, offset: 2240},
												name: "DeclarationSequence",
											},
											&ruleRefExpr{
												pos:  position{line: 86, col: 80, offset: 2262},
												name: "Declaration",
											},
										},
									},
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 88, col: 5, offset: 2335},
						name: "Declaration",
					},
				},
			},
		},
		{
			name: "Declaration",
			pos:  position{line: 90, col: 1, offset: 2348},
			expr: &actionExpr{
				pos: position{line: 90, col: 16, offset: 2363},
				run: (*parser).callonDeclaration1,
				expr: &seqExpr{
					pos: position{line: 90, col: 16, offset: 2363},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 90, col: 16, offset: 2363},
							label: "declOp",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 23, offset: 2370},
								name: "DeclOp",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 30, offset: 2377},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 90, col: 33, offset: 2380},
							label: "initSeq",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 41, offset: 2388},
								name: "InitSeq",
							},
						},
					},
				},
			},
		},
		{
			name: "InitSeq",
			pos:  position{line: 94, col: 1, offset: 2450},
			expr: &actionExpr{
				pos: position{line: 94, col: 12, offset: 2461},
				run: (*parser).callonInitSeq1,
				expr: &seqExpr{
					pos: position{line: 94, col: 12, offset: 2461},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 12, offset: 2461},
							label: "sInit",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 18, offset: 2467},
								name: "SingleInit",
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 29, offset: 2478},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 94, col: 34, offset: 2483},
								expr: &seqExpr{
									pos: position{line: 94, col: 36, offset: 2485},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 94, col: 36, offset: 2485},
											name: "_",
										},
										&litMatcher{
											pos:        position{line: 94, col: 38, offset: 2487},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 42, offset: 2491},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 94, col: 44, offset: 2493},
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
			pos:  position{line: 98, col: 1, offset: 2560},
			expr: &actionExpr{
				pos: position{line: 98, col: 11, offset: 2570},
				run: (*parser).callonDeclOp1,
				expr: &choiceExpr{
					pos: position{line: 98, col: 12, offset: 2571},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 12, offset: 2571},
							val:        "var",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 98, col: 20, offset: 2579},
							val:        "const",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Expression",
			pos:  position{line: 102, col: 1, offset: 2624},
			expr: &choiceExpr{
				pos: position{line: 102, col: 15, offset: 2638},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 102, col: 15, offset: 2638},
						name: "ArithExpr",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 15, offset: 2664},
						name: "BoolExp",
					},
				},
			},
		},
		{
			name: "ArithExpr",
			pos:  position{line: 105, col: 1, offset: 2673},
			expr: &actionExpr{
				pos: position{line: 105, col: 14, offset: 2686},
				run: (*parser).callonArithExpr1,
				expr: &seqExpr{
					pos: position{line: 105, col: 14, offset: 2686},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 105, col: 14, offset: 2686},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 20, offset: 2692},
								name: "Term",
							},
						},
						&labeledExpr{
							pos:   position{line: 105, col: 25, offset: 2697},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 105, col: 30, offset: 2702},
								expr: &seqExpr{
									pos: position{line: 105, col: 32, offset: 2704},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 105, col: 32, offset: 2704},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 105, col: 34, offset: 2706},
											name: "AddOp",
										},
										&ruleRefExpr{
											pos:  position{line: 105, col: 40, offset: 2712},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 105, col: 42, offset: 2714},
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
			pos:  position{line: 109, col: 1, offset: 2770},
			expr: &actionExpr{
				pos: position{line: 109, col: 9, offset: 2778},
				run: (*parser).callonTerm1,
				expr: &seqExpr{
					pos: position{line: 109, col: 9, offset: 2778},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 109, col: 9, offset: 2778},
							label: "first",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 15, offset: 2784},
								name: "Factor",
							},
						},
						&labeledExpr{
							pos:   position{line: 109, col: 22, offset: 2791},
							label: "rest",
							expr: &zeroOrMoreExpr{
								pos: position{line: 109, col: 27, offset: 2796},
								expr: &seqExpr{
									pos: position{line: 109, col: 29, offset: 2798},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 109, col: 29, offset: 2798},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 109, col: 31, offset: 2800},
											name: "MulOp",
										},
										&ruleRefExpr{
											pos:  position{line: 109, col: 37, offset: 2806},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 109, col: 39, offset: 2808},
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
			pos:  position{line: 113, col: 1, offset: 2866},
			expr: &choiceExpr{
				pos: position{line: 113, col: 11, offset: 2876},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 113, col: 11, offset: 2876},
						run: (*parser).callonFactor2,
						expr: &seqExpr{
							pos: position{line: 113, col: 11, offset: 2876},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 113, col: 11, offset: 2876},
									val:        "(",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 113, col: 15, offset: 2880},
									label: "expr",
									expr: &ruleRefExpr{
										pos:  position{line: 113, col: 20, offset: 2885},
										name: "ArithExpr",
									},
								},
								&litMatcher{
									pos:        position{line: 113, col: 30, offset: 2895},
									val:        ")",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 115, col: 5, offset: 2926},
						run: (*parser).callonFactor8,
						expr: &labeledExpr{
							pos:   position{line: 115, col: 5, offset: 2926},
							label: "integer",
							expr: &ruleRefExpr{
								pos:  position{line: 115, col: 13, offset: 2934},
								name: "Integer",
							},
						},
					},
					&actionExpr{
						pos: position{line: 117, col: 5, offset: 2972},
						run: (*parser).callonFactor11,
						expr: &labeledExpr{
							pos:   position{line: 117, col: 5, offset: 2972},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 117, col: 9, offset: 2976},
								name: "Identifier",
							},
						},
					},
				},
			},
		},
		{
			name: "AddOp",
			pos:  position{line: 121, col: 1, offset: 3011},
			expr: &actionExpr{
				pos: position{line: 121, col: 10, offset: 3020},
				run: (*parser).callonAddOp1,
				expr: &choiceExpr{
					pos: position{line: 121, col: 11, offset: 3021},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 121, col: 11, offset: 3021},
							val:        "+",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 121, col: 17, offset: 3027},
							val:        "-",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MulOp",
			pos:  position{line: 125, col: 1, offset: 3069},
			expr: &actionExpr{
				pos: position{line: 125, col: 10, offset: 3078},
				run: (*parser).callonMulOp1,
				expr: &choiceExpr{
					pos: position{line: 125, col: 12, offset: 3080},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 125, col: 12, offset: 3080},
							val:        "*",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 125, col: 18, offset: 3086},
							val:        "/",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoolExp",
			pos:  position{line: 129, col: 1, offset: 3128},
			expr: &choiceExpr{
				pos: position{line: 129, col: 12, offset: 3139},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 129, col: 12, offset: 3139},
						run: (*parser).callonBoolExp2,
						expr: &labeledExpr{
							pos:   position{line: 129, col: 12, offset: 3139},
							label: "keyword",
							expr: &ruleRefExpr{
								pos:  position{line: 129, col: 21, offset: 3148},
								name: "KeywordBoolExp",
							},
						},
					},
					&actionExpr{
						pos: position{line: 132, col: 3, offset: 3213},
						run: (*parser).callonBoolExp5,
						expr: &seqExpr{
							pos: position{line: 132, col: 3, offset: 3213},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 132, col: 3, offset: 3213},
									label: "unBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 13, offset: 3223},
										name: "UnaryBoolOp",
									},
								},
								&labeledExpr{
									pos:   position{line: 132, col: 25, offset: 3235},
									label: "boolExp",
									expr: &ruleRefExpr{
										pos:  position{line: 132, col: 34, offset: 3244},
										name: "BoolExp",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 135, col: 3, offset: 3310},
						run: (*parser).callonBoolExp11,
						expr: &seqExpr{
							pos: position{line: 135, col: 3, offset: 3310},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 135, col: 3, offset: 3310},
									label: "id1",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 8, offset: 3315},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 135, col: 19, offset: 3326},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 135, col: 21, offset: 3328},
									label: "binBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 32, offset: 3339},
										name: "BinaryBoolOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 135, col: 45, offset: 3352},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 135, col: 47, offset: 3354},
									label: "id2",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 52, offset: 3359},
										name: "Identifier",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 138, col: 3, offset: 3432},
						run: (*parser).callonBoolExp21,
						expr: &seqExpr{
							pos: position{line: 138, col: 3, offset: 3432},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 138, col: 3, offset: 3432},
									label: "id1",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 8, offset: 3437},
										name: "Identifier",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 19, offset: 3448},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 21, offset: 3450},
									label: "binBoolOp",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 32, offset: 3461},
										name: "BinaryBoolOp",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 138, col: 45, offset: 3474},
									name: "_",
								},
								&labeledExpr{
									pos:   position{line: 138, col: 47, offset: 3476},
									label: "id2",
									expr: &ruleRefExpr{
										pos:  position{line: 138, col: 51, offset: 3480},
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
			pos:  position{line: 142, col: 1, offset: 3551},
			expr: &actionExpr{
				pos: position{line: 142, col: 19, offset: 3569},
				run: (*parser).callonKeywordBoolExp1,
				expr: &choiceExpr{
					pos: position{line: 142, col: 20, offset: 3570},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 142, col: 20, offset: 3570},
							val:        "true",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 142, col: 29, offset: 3579},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "UnaryBoolOp",
			pos:  position{line: 146, col: 1, offset: 3624},
			expr: &actionExpr{
				pos: position{line: 146, col: 16, offset: 3639},
				run: (*parser).callonUnaryBoolOp1,
				expr: &litMatcher{
					pos:        position{line: 146, col: 16, offset: 3639},
					val:        "~",
					ignoreCase: false,
				},
			},
		},
		{
			name: "BinaryBoolOp",
			pos:  position{line: 150, col: 1, offset: 3678},
			expr: &actionExpr{
				pos: position{line: 150, col: 17, offset: 3694},
				run: (*parser).callonBinaryBoolOp1,
				expr: &choiceExpr{
					pos: position{line: 150, col: 19, offset: 3696},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 150, col: 19, offset: 3696},
							val:        "==",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 26, offset: 3703},
							val:        "<=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 33, offset: 3710},
							val:        ">=",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 40, offset: 3717},
							val:        "<",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 46, offset: 3723},
							val:        ">",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 52, offset: 3729},
							val:        "/\\",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 150, col: 60, offset: 3737},
							val:        "\\/",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Identifier",
			pos:  position{line: 154, col: 1, offset: 3780},
			expr: &actionExpr{
				pos: position{line: 154, col: 15, offset: 3794},
				run: (*parser).callonIdentifier1,
				expr: &seqExpr{
					pos: position{line: 154, col: 15, offset: 3794},
					exprs: []interface{}{
						&charClassMatcher{
							pos:        position{line: 154, col: 15, offset: 3794},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 154, col: 23, offset: 3802},
							expr: &charClassMatcher{
								pos:        position{line: 154, col: 23, offset: 3802},
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
			pos:  position{line: 158, col: 1, offset: 3850},
			expr: &actionExpr{
				pos: position{line: 158, col: 12, offset: 3861},
				run: (*parser).callonInteger1,
				expr: &seqExpr{
					pos: position{line: 158, col: 12, offset: 3861},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 158, col: 12, offset: 3861},
							expr: &litMatcher{
								pos:        position{line: 158, col: 12, offset: 3861},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 158, col: 17, offset: 3866},
							expr: &charClassMatcher{
								pos:        position{line: 158, col: 17, offset: 3866},
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
			name: "Real",
			pos:  position{line: 162, col: 1, offset: 3909},
			expr: &actionExpr{
				pos: position{line: 162, col: 9, offset: 3917},
				run: (*parser).callonReal1,
				expr: &seqExpr{
					pos: position{line: 162, col: 9, offset: 3917},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 162, col: 9, offset: 3917},
							expr: &choiceExpr{
								pos: position{line: 162, col: 10, offset: 3918},
								alternatives: []interface{}{
									&litMatcher{
										pos:        position{line: 162, col: 10, offset: 3918},
										val:        "+",
										ignoreCase: false,
									},
									&litMatcher{
										pos:        position{line: 162, col: 16, offset: 3924},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 162, col: 22, offset: 3930},
							expr: &charClassMatcher{
								pos:        position{line: 162, col: 22, offset: 3930},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&litMatcher{
							pos:        position{line: 162, col: 29, offset: 3937},
							val:        ".",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 162, col: 33, offset: 3941},
							expr: &charClassMatcher{
								pos:        position{line: 162, col: 33, offset: 3941},
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
			name: "Text",
			pos:  position{line: 166, col: 1, offset: 3984},
			expr: &actionExpr{
				pos: position{line: 166, col: 8, offset: 3993},
				run: (*parser).callonText1,
				expr: &seqExpr{
					pos: position{line: 166, col: 8, offset: 3993},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 166, col: 8, offset: 3993},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 166, col: 12, offset: 3997},
							expr: &choiceExpr{
								pos: position{line: 166, col: 14, offset: 3999},
								alternatives: []interface{}{
									&seqExpr{
										pos: position{line: 166, col: 14, offset: 3999},
										exprs: []interface{}{
											&notExpr{
												pos: position{line: 166, col: 14, offset: 3999},
												expr: &ruleRefExpr{
													pos:  position{line: 166, col: 15, offset: 4000},
													name: "EscapedChar",
												},
											},
											&anyMatcher{
												line: 166, col: 27, offset: 4012,
											},
										},
									},
									&seqExpr{
										pos: position{line: 166, col: 31, offset: 4016},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 166, col: 31, offset: 4016},
												val:        "\\",
												ignoreCase: false,
											},
											&ruleRefExpr{
												pos:  position{line: 166, col: 36, offset: 4021},
												name: "EscapeSequence",
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 166, col: 54, offset: 4039},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedChar",
			pos:  position{line: 170, col: 1, offset: 4079},
			expr: &charClassMatcher{
				pos:        position{line: 170, col: 15, offset: 4095},
				val:        "[\\x00-\\x1f\"\\\\]",
				chars:      []rune{'"', '\\'},
				ranges:     []rune{'\x00', '\x1f'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EscapeSequence",
			pos:  position{line: 172, col: 1, offset: 4111},
			expr: &choiceExpr{
				pos: position{line: 172, col: 18, offset: 4130},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 172, col: 18, offset: 4130},
						name: "SingleCharEscape",
					},
					&ruleRefExpr{
						pos:  position{line: 172, col: 37, offset: 4149},
						name: "UnicodeEscape",
					},
				},
			},
		},
		{
			name: "SingleCharEscape",
			pos:  position{line: 174, col: 1, offset: 4164},
			expr: &charClassMatcher{
				pos:        position{line: 174, col: 20, offset: 4185},
				val:        "[\"\\\\/bfnrt]",
				chars:      []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "UnicodeEscape",
			pos:  position{line: 176, col: 1, offset: 4198},
			expr: &seqExpr{
				pos: position{line: 176, col: 17, offset: 4216},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 176, col: 17, offset: 4216},
						val:        "u",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 176, col: 21, offset: 4220},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 176, col: 30, offset: 4229},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 176, col: 39, offset: 4238},
						name: "HexDigit",
					},
					&ruleRefExpr{
						pos:  position{line: 176, col: 48, offset: 4247},
						name: "HexDigit",
					},
				},
			},
		},
		{
			name: "HexDigit",
			pos:  position{line: 178, col: 1, offset: 4257},
			expr: &charClassMatcher{
				pos:        position{line: 178, col: 12, offset: 4270},
				val:        "[0-9a-f]i",
				ranges:     []rune{'0', '9', 'a', 'f'},
				ignoreCase: true,
				inverted:   false,
			},
		},
		{
			name:        "_",
			displayName: "\"whitespace\"",
			pos:         position{line: 181, col: 1, offset: 4282},
			expr: &zeroOrMoreExpr{
				pos: position{line: 181, col: 19, offset: 4300},
				expr: &charClassMatcher{
					pos:        position{line: 181, col: 19, offset: 4300},
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
			pos:         position{line: 182, col: 1, offset: 4311},
			expr: &oneOrMoreExpr{
				pos: position{line: 182, col: 31, offset: 4341},
				expr: &charClassMatcher{
					pos:        position{line: 182, col: 31, offset: 4341},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 184, col: 1, offset: 4353},
			expr: &notExpr{
				pos: position{line: 184, col: 8, offset: 4360},
				expr: &anyMatcher{
					line: 184, col: 9, offset: 4361,
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

func (c *current) onClauses1(variable, constant, init, procs, calls interface{}) (interface{}, error) {
	return constructClauses(variable, constant, init, procs, calls), nil
}

func (p *parser) callonClauses1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onClauses1(stack["variable"], stack["constant"], stack["init"], stack["procs"], stack["calls"])
}

func (c *current) onProcedure1(id, formals, blk interface{}) (interface{}, error) {
	return constructProcedure(id, formals, blk), nil
}

func (p *parser) callonProcedure1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onProcedure1(stack["id"], stack["formals"], stack["blk"])
}

func (c *current) onFormals2(first, rest interface{}) (interface{}, error) {
	return constructFormals(first, rest), nil
}

func (p *parser) callonFormals2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFormals2(stack["first"], stack["rest"])
}

func (c *current) onFormals12() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonFormals12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFormals12()
}

func (c *current) onCall1(id, actuals interface{}) (interface{}, error) {
	return constructCall(id, actuals), nil
}

func (p *parser) callonCall1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCall1(stack["id"], stack["actuals"])
}

func (c *current) onActuals2(first, rest interface{}) (interface{}, error) {
	return constructActuals(first, rest), nil
}

func (p *parser) callonActuals2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onActuals2(stack["first"], stack["rest"])
}

func (c *current) onActuals12() (interface{}, error) {
	return nil, nil
}

func (p *parser) callonActuals12() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onActuals12()
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

func (c *current) onDeclarationSequence2(first, rest interface{}) (interface{}, error) {
	return evalDeclarationSequence(first, rest), nil
}

func (p *parser) callonDeclarationSequence2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDeclarationSequence2(stack["first"], stack["rest"])
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

func (c *current) onReal1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonReal1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReal1()
}

func (c *current) onText1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonText1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onText1()
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
