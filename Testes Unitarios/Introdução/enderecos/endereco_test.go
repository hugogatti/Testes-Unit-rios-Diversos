package enderecos

import "testing"

//TESTES UNITARIO BASICOS

func TestTipoDeEndereco(t *testing.T) {
	//FOMATO OBRIGADORIO o nome do arquivo tem que ter "_test" e o nome da func tem que ser neste formato: TestNomeDaSuaFunçao
	enderecoParaTeste := "Rua Manuel Justiniano Quintão"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}