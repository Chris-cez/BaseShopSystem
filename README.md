# Sistema de vendas base
Versão básica e open-source do código usado na criação de um sistema de uma loja real. A versão completa com os dados especificos do projeto será criada e vendida de forma privada. Este documento descreve a estrutura de dados utilizada para representar uma Nota Fiscal Eletrônica (NF-e) em um sistema de vendas.

## Tabelas do Banco de Dados

A estrutura é baseada em múltiplas tabelas relacionadas para garantir a integridade e eficiência dos dados. Pensada para posteriormente ser usada com GORM.

**1. `Empresa` (Dados da Empresa - Hardcoded)**

| Campo               | Tipo de Dado  | Descrição                                        |
|---------------------|---------------|--------------------------------------------------|
| `Nome`              | `string`      | Nome da empresa                                  |
| `CNPJ`              | `string`      | Número do CNPJ da empresa                        |
| `InscricaoEstadual` | `string`      | Inscrição Estadual da empresa                    |


**2. `Endereco` (Dados de Endereço)**

| Campo        | Tipo de Dado  | Descrição                                 |
|--------------|---------------|-------------------------------------------|
| `ID`         | `uint`        | Chave primária                            |
| `Logradouro` | `string`      | Rua, Avenida, etc.                        |
| `Numero`     | `string`      | Número do endereço                        |
| `Complemento`| `string`      | Complemento do endereço                   |
| `Bairro`     | `string`      | Bairro                                    |
| `Municipio`  | `string`      | Município                                 |
| `UF`         | `string`      | Unidade Federativa                        |
| `CEP`        | `string`      | Código de Endereçamento Postal            |


**3. `Cliente` (Dados do Cliente)**

| Campo             | Tipo de Dado  | Descrição                                        |
|-------------------|---------------|--------------------------------------------------|
| `ID`              | `uint`        | Chave primária                                   |
| `Nome`            | `string`      | Nome do cliente                                  |
| `CP`              | `string`      | CPF do cliente (ou CNPJ)                         |
| `EnderecoID`      | `uint`        | Chave estrangeira (Endereco.ID)                  |
| `Telefone`        | `string`      | Telefone do cliente                              |
| `Email`           | `string`      | Email do cliente                                 |


**4. `ClasseProduto` (Classificação do Produto)**

| Campo             | Tipo de Dado   | Descrição                                          |
|-------------------|----------------|----------------------------------------------------|
| `ID`              | `uint`         | Chave primária                                     |
| `Nome`            | `string`       | Nome da classe do produto                          |
| `Descricao`       | `string`       | Descrição da classe do produto                     |
| `NCM`             | `string`       | Código NCM da classe do produto                    |
| `Tributacoes`     | `[]Tributacao` | Array de tributações aplicadas à classe do produto |


**5. `Produto` (Dados do Produto)**

| Campo             | Tipo de Dado  | Descrição                                      |
|-------------------|---------------|------------------------------------------------|
| `ID`              | `uint`        | Chave primária                                 |
| `GTIN`            | `string`      | GTIN do produto                                |
| `Nome`            | `string`      | Nome do produto                                |
| `Descricao`       | `string`      | Descrição do produto                           |
| `ClasseProdutoID` | `uint`        | Chave estrangeira (ClasseProduto.ID)           |
| `ValorUnitario`   | `float64`     | Valor unitário do produto                      |
| `Unidade`         | `string`      | Unidade de medida (ex: UN, KG, M)              |
| `Estoque`         | `int`         | Quantidade em estoque                          |


**6. `Tributacao` (Dados de Tributação)**

| Campo             | Tipo de Dado | Descrição                                      |
|-------------------|--------------|------------------------------------------------|
| `ID`              | `uint`       | Chave primária                                 |
| `Nome`            | `string`     | Nome do tributo (ex: ICMS, IPI, PIS, COFINS)   |
| `Aliquota`        | `float64`    | Alíquota do tributo                            |
| `TipoTributo`     | `string`     | Tipo de tributo                                |



**7. `ItemNotaFiscal` (Itens da Nota Fiscal)**

| Campo             | Tipo de Dado  | Descrição                                        |
|-------------------|---------------|--------------------------------------------------|
| `ID`              | `uint`        | Chave primária                                   |
| `NotaFiscalID`    | `uint`        | Chave estrangeira (NotaFiscal.ID)                |
| `ProdutoID`       | `uint`        | Chave estrangeira (Produto.ID)                   |
| `Quantidade`      | `float64`     | Quantidade do produto                            |
| `ValorUnitario`   | `float64`     | Valor unitário do produto na nota                |
| `ValorTotal`      | `float64`     | Valor total do item (Quantidade * ValorUnitario) |


**8. `NotaFiscal` (Dados da Nota Fiscal)**

| Campo             | Tipo de Dado  | Descrição                                       |
|-------------------|---------------|-------------------------------------------------|
| `ID`              | `gorm.Model`  | Chave primária (GORM)                           |
| `Numero`          | `string`      | Número da NF-e                                  |
| `ClienteID`       | `uint`        | Chave estrangeira (Cliente.ID)                  |
| `ValorTotal`      | `float64`     | Valor total da nota fiscal                      |
| `FormaPagamentoID`| `uint`        | Chave estrangeira (FormaPagamento.ID)           |
| `Desconto`        | `float64`     | Valor do desconto aplicado                      |
| `Observacao`      | `string`      | Observações adicionais                          |
| `ChaveAcesso`     | `string`      | Chave de acesso da NF-e (após autorização)      |


**9. `FormaPagamento` (Formas de Pagamento)**

| Campo             | Tipo de Dado | Descrição                                         |
|-------------------|--------------|---------------------------------------------------|
| `ID`              | `uint`       | Chave primária                                    |
| `Nome`            | `string`     | Nome da forma de pagamento (ex: Dinheiro, Cartão) |


Este é um resumo. Uma NF-e real contém muitos mais detalhes e campos específicos dependendo da legislação e do tipo de operação. 
