// TESTES UNITARIO COM DIVERSOS CENARIOS
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste {
		{ "Rua ABC", "Rua" },
		{ "Avenida Paulista", "Avenida" },
		{ "Rodovia dos Imigrantes", "Rodovia" },
		{ "Praça do Cruzeiro", "Tipo Inválido" },
		{ "Estrada dos Loucos", "Estrada" },
		{ "RUA DOS BOBOS", "Rua" },
		{ "AVENIDA REBOUÇAS", "Avenida" },
		{ "", "Tipo Inválido" },
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)			
		}
	}
}