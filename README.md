# Base Shop System
Versão básica e open-source do código usado na criação de um sistema de uma loja real. A versão completa com os dados especificos do projeto será criada e vendida de forma privada. Se tiver interesse na versão paga, tratar com o criador do repositorio original.

## Banco de Dados
Versão pensada para o **GORM**.

**Product**
| Name | Type |
| :---: | :---: |
| Id | string | Primary key
| Name | string |
| Price | float32 |
| NCM | string |
| UM | string |
| Class | uint8 | Foreign key
| Description | string |

**Client**
| Name | Type |
| :---: | :---: |
| Id | string | Primary key
| Name | string |
| CPF | int32 | 
| Address | uint16 | Foreign key


`
type User struct {
	gorm.Model
	Name string
	CPF  uint32
	Id_address uint16
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
}

type Product struct {
	gorm.Model
	Code  string
	Price float32
	Name  string
	NCM   string
	UM    string
	Description string
	Id_Class uint8
}

type Class struct {
	gorm.Model
	Name string
	Description string
}
`
