package parser

import (
	"jeisaraja/interpreter/ast"
	"jeisaraja/interpreter/lexer"
	"jeisaraja/interpreter/token"
)

type Parser struct {
  l *lexer.Lexer
  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser{
  p := &Parser{l: l}
  
  return p
}

func (p *Parser) nextToken(){
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
